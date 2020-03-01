package gohome

type Event struct{}

type EventChannel <-chan Event

type ListenFunc func() EventChannel

func NewEventSource() (chan<- Event, func() EventChannel) {
	listen := make(chan chan Event)
	src := make(chan Event, 256)

	dispatch := func() {
		var dests []chan Event
		// TODO(dape): add shutdown
		for {
			select {
			case e := <-src:
				for _, dest := range dests {
					// TODO(dape): don't block on send
					dest <- e
				}

			case dest := <-listen:
				switch {
				case dest != nil:
					dests = append(dests, dest)
				}
			}
		}
	}

	go dispatch()

	return src, func() EventChannel {
		dest := make(chan Event, 256)
		listen <- dest
		return dest
	}
}
