package repository

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
)

// SaveInput is the input for the Save method.
type SaveInput struct {
	Diet Diet
}

// AllInput is the input for the All method.
type AllInput struct{}

// FindInput is the input for the Find method.
type FindInput struct {
	ID value.ID
}

// DeleteInput is the input for the Delete method.
type DeleteInput struct {
	ID value.ID
}

// UpdateInput is the input for the Update method.
type UpdateInput struct {
	Diet Diet
}

// Diet is the interface that wraps the basic Diet methods.
type Diet interface {
	// Save saves the given diet.
	Save(ctx context.Context, input SaveInput) error

	// All returns all diets.
	All(ctx context.Context, input AllInput) ([]*Diet, error)

	// Find returns the diet with the given id.
	Find(ctx context.Context, input FindInput) (*Diet, error)

	// Delete deletes the diet with the given id.
	Delete(ctx context.Context, input DeleteInput) error

	// Update updates the given diet.
	Update(ctx context.Context, input UpdateInput) error
}
