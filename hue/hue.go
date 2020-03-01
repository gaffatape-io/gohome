package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gaffatape-io/gohome"
	errs "github.com/gaffatape-io/gopherrs"
	"k8s.io/klog"
)

type Bridge struct {
	URL      string
	Username string
	Client   *http.Client
}

func newJSONRequest(method string, suffix string, body interface{}) (*http.Request, error) {
	// TODO(dape): add code to make normal json a 'one-liner'

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, suffix, buf)
}

type errorDetails struct {
	Address     string `json:address`
	Description string `json:description`
	Type        int    `json:type`
}

func (e *errorDetails) Error() string {
	return fmt.Sprint("Hue Error; type:", e.Type, " Desc:", e.Description, " Address:", e.Address)
}

type createUserContainer struct {
	Success createUserResponse `json:success`
	Error   errorDetails       `json:error`
}

func (b *Bridge) httpURL(suffix string) string {
	return fmt.Sprint(b.URL, "/api/", b.Username, suffix)
}

func (b *Bridge) httpCheck(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return nil, errs.Wrap(err)
	}

	if resp.StatusCode != http.StatusOK {
		// TODO(dape): implement me (or add to gopherrs)
		panic("translate error")
	}

	return resp, err
}

type request interface {
	// toHTTP returns the corresponding HTTP request.
	toHTTP() (*http.Request, error)
}

type response interface {
	// fromHTTP populates from response or returns an error.
	// Method is never called for a non successul HTTP status code.
	// JSON data from body is already marshaled into this instance
	// before method is called.
	fromHTTP(*http.Response) error
}

func (b *Bridge) roundtrip(resp response, req request) error {
	httpreq, err := req.toHTTP()
	if err != nil {
		return err
	}

	// TODO(dape): move parsing to bridge object creation
	bu, err := url.Parse(fmt.Sprint(b.URL, "/api/", b.Username))
	if err != nil {
		return errs.Internal(err)
	}

	httpreq.URL = httpreq.URL.ResolveReference(bu)
	klog.Info("roundtrip:", req, httpreq)
	httpresp, err := b.Client.Do(httpreq)
	if err != nil {
		return errs.Internal(err)
	}

	if httpresp.StatusCode != http.StatusOK {
		return errs.Internal(fmt.Errorf("http error:%d", httpresp.StatusCode))
	}

	return resp.fromHTTP(httpresp)
}

type Bulb struct {
	Name string
}

func (b *Bulb) On() {
}

func (b Bulb) Off() {
}

type BulbState struct {
	On int
	D  time.Duration
}

func (b *Bulb) Update(state BulbState) {
}

type MotionSensor struct {
	ID     string
	sink   chan<- gohome.Event
	listen gohome.ListenFunc
}

func NewMotionSensor(id string) *MotionSensor {
	src, listen := gohome.NewEventSource()
	m := &MotionSensor{id, src, listen}
	return m
}
