package event

import (
	"fmt"
)

// Event is the interface that wraps the basic methods of an event.
type Event[Data any] interface {
	String() string
	Action() Action
	Name() Name
}

// eventImpl is the type of the event.
type eventImpl[Data any] struct {
	name   Name
	action Action
}

// New returns a new event.
func New[Data any](name Name, action Action) Event[Data] {
	return &eventImpl[Data]{
		name:   name,
		action: action,
	}
}

// String returns the string representation of the event.
func (e eventImpl[Data]) String() string {
	return fmt.Sprintf("%s.%s", e.name, e.action)
}

// Action returns the action of the event.
func (e eventImpl[Data]) Action() Action {
	return e.action
}

// Name returns the name of the event.
func (e eventImpl[Data]) Name() Name {
	return e.name
}
