package events

import (
	"reflect"
	"time"
)

// Returns a synchronous event emitter.
func Subscribe(evsw EventSwitch, listenerID string) <-chan Event {
	ch := make(chan Event, 0) // synchronous
	return SubscribeOn(evsw, listenerID, ch)
}

// Like Subscribe, but lets the caller construct a channel.  If the capacity of
// the provided channel is 0, it will be called synchronously; otherwise, it
// will drop when the capacity is reached and a select doesn't immediately
// send.
func SubscribeOn(evsw EventSwitch, listenerID string, ch chan Event) <-chan Event {
	return SubscribeFilteredOn(evsw, listenerID, nil, ch)
}

func SubscribeToEvent(evsw EventSwitch, listenerID string, protoevent Event) <-chan Event {
	ch := make(chan Event, 0) // synchronous
	return SubscribeToEventOn(evsw, listenerID, protoevent, ch)
}

func SubscribeToEventOn(evsw EventSwitch, listenerID string, protoevent Event, ch chan Event) <-chan Event {
	rt := reflect.TypeOf(protoevent)
	return SubscribeFilteredOn(evsw, listenerID, &TypeFilterer{Type: rt}, ch)
}

type EventFilter func(Event) bool

func SubscribeFiltered(evsw EventSwitch, listenerID string, filterer Filterer) <-chan Event {
	ch := make(chan Event, 0)
	return SubscribeFilteredOn(evsw, listenerID, filterer, ch)
}

func SubscribeFilteredOn(evsw EventSwitch, listenerID string, filterer Filterer, ch chan Event) <-chan Event {
	listener := NewEventListener(evsw, listenerID, ch, filterer, 10*time.Second)
	evsw.AddListener(listener)
	return ch
}
