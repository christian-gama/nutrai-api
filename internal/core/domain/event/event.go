package event

import (
	"fmt"
)

// Event is the interface that wraps the basic methods of an event.
type Event interface {
	String() string
	Action() Action
	Name() Name
}

// eventImpl is the type of the event.
type eventImpl struct {
	name   Name
	action Action
}

// New returns a new event.
func New(name Name, action Action) Event {
	return &eventImpl{
		name:   name,
		action: action,
	}
}

// String returns the string representation of the event.
func (e eventImpl) String() string {
	return fmt.Sprintf("%s.%s", e.name, e.action)
}

// Action returns the action of the event.
func (e eventImpl) Action() Action {
	return e.action
}

// Name returns the name of the event.
func (e eventImpl) Name() Name {
	return e.name
}
