package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gaffatape-io/gohome"
	"github.com/gaffatape-io/gohome/hue"
	"k8s.io/klog"
)

var (
	hueBridgeURL      = os.Getenv("HUE_BRIDGE_URL")
	hueBridgeApp      = os.Getenv("HUE_BRIDGE_APP")
	hueBridgeUser     = os.Getenv("HUE_BRIDGE_USER")
	hueBridgeUsername = os.Getenv("HUE_BRIDGE_USERNAME")
)

type Hallway struct{}

type HubertLaurin1 struct {
	Hallway Hallway
}

func (h *HubertLaurin1) Run() error {
	return nil
}

func main() {
	klog.Info("bridge-url: ", hueBridgeURL)
	klog.Info("bridge-app: ", hueBridgeApp)
	klog.Info("bridge-user: ", hueBridgeUser)
	klog.Info("bridge-username: ", hueBridgeUsername)

	bridge := &hue.Bridge{hueBridgeURL, hueBridgeUsername, &http.Client{}}
	if bridge.Username == "" {
		err := bridge.Link(hueBridgeApp, hueBridgeUser, 3*time.Minute)
		if err != nil {
			klog.Fatalf("link failed; err:%+v", err)
		}
		klog.Infof("link succeed; username:%s", bridge.Username)
	}

	klog.Info(bridge.Lights())

	env := gohome.NewEnvironment()
	err := env.Run(":8080")
	if err != nil {
	}
}
