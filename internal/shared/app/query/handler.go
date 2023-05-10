package query

import "context"

// Handler is the interface that wraps the Handle method.
type Handler[Input any, Output any] interface {
	// Handle processes the given input and returns the output and/or an error.
	// The context can be used to pass any information that may be needed.
	Handle(ctx context.Context, input Input) (Output, error)
}
