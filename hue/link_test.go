package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFormatBridgeDeviceType(t *testing.T) {
	tests := []struct {
		app    string
		user   string
		result string
	}{
		{"hello-world", "dape", "hello-world#dape"},
		{"", "", ""},
	}

	for _, tc := range tests {
		res, err := formatBridgeDeviceType(tc.app, tc.user)
		t.Log(tc, res, err)
		if tc.result != res {
			t.Fatal()
		}

		if tc.result == "" && err == nil {
			t.Fatal()
		}
	}
}

type linkResponse struct {
	Username string `json:"username`
}

func TestLinkBridge(t *testing.T) {
	tests := []struct {
		attemptsRequired int
		username         string
	}{
		{1, "username1"},
		{3, "username3"},
	}

	for _, tc := range tests {
		t.Log(tc)
		attempt := 0

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			attempt++
			var tuple createUserResponseTuple
			if attempt == tc.attemptsRequired {
				tuple = createUserResponseTuple{Success: createUserResponse{Username: fmt.Sprint("username", attempt)}}
			} else {
				tuple = createUserResponseTuple{Error: &errorDetails{Type: 101}}
			}

			bytes, _ := json.Marshal([]createUserResponseTuple{tuple})
			w.Write(bytes)
		}))
		defer srv.Close()

		b := &Bridge{srv.URL, "", &http.Client{}}
		sleepCount := 0
		tcur := time.Time{}
		funcs := linkFuncs{
			func(duration time.Duration) {
				sleepCount++
				t.Log("sleep:", duration)
			},

			func() time.Time {
				tcur = tcur.Add(1 * time.Second)
				return tcur
			},
		}

		err := b.link("test", t.Name(), 3*time.Second, funcs)
		t.Log(err)
		if err != nil {
			t.Fatal()
		}

		if sleepCount+1 != tc.attemptsRequired {
			t.Fatal(sleepCount)
		}

		t.Log(b)
		if tc.username != b.Username {
			t.Fatal()
		}
	}

}
