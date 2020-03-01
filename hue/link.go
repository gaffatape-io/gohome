package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	errs "github.com/gaffatape-io/gopherrs"
	"k8s.io/klog"
)

func formatBridgeDeviceType(app, user string) (string, error) {
	const maxDeviceType = 40
	deviceType := fmt.Sprintf("%s#%s", app, user)
	if len(deviceType) > maxDeviceType || len(app) == 0 || len(user) == 0 {
		return "", errs.InvalidArgumentf(nil, "app:(%d/20) user:(%d/19)", len(app), len(user))
	}

	return deviceType, nil
}

type createUserRequest struct {
	App  string
	User string
}

func (c *createUserRequest) toHTTP() (*http.Request, error) {
	dt, err := formatBridgeDeviceType(c.App, c.User)
	if err != nil {
		return nil, err
	}

	req := struct {
		DeviceType string `json:"devicetype"`
	}{
		dt,
	}

	return newJSONRequest(http.MethodPost, "", req)
}

type createUserResponse struct {
	Username string `json:username`
}

type createUserResponseTuple struct {
	Success createUserResponse `json:"success"`
	Error   *errorDetails      `json:"error"`
}

func (c *createUserResponse) fromHTTP(resp *http.Response) error {
	tuples := []createUserResponseTuple{}
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&tuples)
	if err != nil {
		return errs.Internal(err)
	}

	if len(tuples) == 0 {
		// TODO(dape): fix this
		panic("doh")
	}

	if tuples[0].Error != nil {
		details := tuples[0].Error
		// TODO(dape): define errors from hue as values
		if details.Type == 101 {
			return errs.FailedPreconditionf(nil, "Link button not pressed")
		}
		return errs.Internal(details)
	}

	*c = tuples[0].Success
	return nil
}

// Link is used to create a whitelisted user in the local bridge.
func (b *Bridge) Link(app, user string, timeout time.Duration) error {
	return b.link(app, user, timeout, linkfuncs)
}

var (
	linkfuncs = linkFuncs{time.Sleep, time.Now}
)

type linkFuncs struct {
	sleep func(time.Duration)
	now   func() time.Time
}

func (b *Bridge) link(app, user string, timeout time.Duration, depfuncs linkFuncs) error {
	deadline := depfuncs.now().Add(timeout)
	req := &createUserRequest{app, user}

	for {
		klog.Info("creating user; press hue bridge link button...")
		now := depfuncs.now()
		if now.After(deadline) {
			return errs.DeadlineExceededf(nil, "failed to link user; deadline:%v - now:%v", deadline, now)
		}

		resp := &createUserResponse{}
		err := b.roundtrip(resp, req)
		if errs.IsFailedPrecondition(err) {
			// Link button not pressed, retry after sleep...
			depfuncs.sleep(500 * time.Millisecond)
			continue
		}

		if err != nil {
			return err
		}

		b.Username = resp.Username
		return nil
	}
}
