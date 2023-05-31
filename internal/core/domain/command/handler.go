package command

import "context"

// Handler is the interface that wraps the Handle method.
type Handler[Input any] interface {
	// Handle processes the given input and returns an error if any.
	// The context can be used to pass any information that may be needed.
	Handle(ctx context.Context, input Input) error
}
