package event

import (
	"reflect"

	"github.com/gopherd/doge/event"
	"github.com/gopherd/doge/event/dispatcher"
)

type Type = reflect.Type

type Event = event.Event[Type]
type Listener = event.Listener[Type]
type ListenerID = event.ListenerID
type Dispatcher = event.Dispatcher[Type]
type BasicDispatcher = dispatcher.Dispatcher[Type]

func Listen[E Event](eventType Type, handler func(E)) Listener {
	return event.Listen(eventType, handler)
}

func TypeOf[E Event](e E) Type {
	// It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	t := reflect.TypeOf(e)
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Pointer, reflect.Slice:
		return t.Elem()
	default:
		return t
	}
}
