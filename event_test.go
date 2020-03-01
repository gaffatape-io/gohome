package gohome

import "testing"

func TestNewEventSource(t *testing.T) {
	c, listen := NewEventSource()

	l0 := listen()
	l1 := listen()

	// If we send this before the call to listen there is a small
	// chance that the message will be dropped, this is due to the
	// latency of the goroutine startup in listen.
	c <- Event{}

	e0recv := false
	e1recv := false

	for i := 0; i < 2; i++ {
		select {
		case <-l0:
			e0recv = true

		case <-l1:
			e1recv = true
		}
	}

	if !e0recv || !e1recv {
		t.Fatal(e0recv, e1recv)
	}
}
