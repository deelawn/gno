package events

import "reflect"

type Filterer interface {
	Filter(Event) bool
}

type TypeFilterer struct {
	Type reflect.Type
}

func (f TypeFilterer) Filter(event Event) bool {
	return reflect.TypeOf(event) == f.Type
}
