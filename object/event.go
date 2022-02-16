package object

import "github.com/gopherd/three/core/event"

type (
	AddedEvent   struct{}
	RemovedEvent struct{}
)

//@mod:final
var (
	AddedEventType   = event.TypeOf[*AddedEvent](nil)
	RemovedEventType = event.TypeOf[*RemovedEvent](nil)
)

func (AddedEvent) Type() event.Type   { return AddedEventType }
func (RemovedEvent) Type() event.Type { return RemovedEventType }

//@mod:final
var (
	addedEvent   = AddedEvent{}
	removedEvent = RemovedEvent{}
)
