package events

import "time"

type Listener interface {
	ID() string
	Listen(Event)
}

type EventListener struct {
	eventSwitch EventSwitch
	id          string
	ch          chan Event
	filterer    Filterer
	timeout     time.Duration
}

func NewEventListener(
	eventSwitch EventSwitch,
	id string,
	ch chan Event,
	filterer Filterer,
	timeout time.Duration,
) *EventListener {
	return &EventListener{
		eventSwitch: eventSwitch,
		id:          id,
		ch:          ch,
		filterer:    filterer,
		timeout:     timeout,
	}
}

func (l EventListener) ID() string {
	return l.id
}

func (l EventListener) Listen(event Event) {
	if l.filterer != nil && !l.filterer.Filter(event) {
		return // filter
	}
	// NOTE: This callback must not block for performance.
	if cap(l.ch) == 0 {
		timeout := l.timeout
	LOOP:
		for {
			select { // sync
			case l.ch <- event:
				break LOOP
			case <-l.eventSwitch.Quit():
				close(l.ch)
				break LOOP
			case <-time.After(timeout):
				// After a minute, print a message for debugging.
				//log.Printf("[WARN] EventSwitch subscriber %v blocked on %v for %v", l.id, event, timeout)
				// Exponentially back off warning messages.
				timeout *= 2
			}
		}
	} else {
		select { // async
		case l.ch <- event:
		case <-l.eventSwitch.Quit():
			close(l.ch)
		default:
		}
	}
}
