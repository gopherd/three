package object

import "github.com/gopherd/three/core/event"

type (
	AddedEvent   struct{}
	RemovedEvent struct{}
)

var (
	AddedEventType   = event.TypeOf[*AddedEvent](nil)
	RemovedEventType = event.TypeOf[*AddedEvent](nil)
)

func (AddedEvent) Type() event.Type   { return AddedEventType }
func (RemovedEvent) Type() event.Type { return RemovedEventType }

var (
	addedEvent   = AddedEvent{}
	removedEvent = RemovedEvent{}
)
