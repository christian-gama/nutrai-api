package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
)

// SaveDietInput is the input for the Save method.
type SaveDietInput struct {
	Diet diet.Diet
}

// AllDietInput is the input for the All method.
type AllDietInput struct{}

// FindDietInput is the input for the Find method.
type FindDietInput struct {
	ID value.ID
}

// DeleteDietInput is the input for the Delete method.
type DeleteDietInput struct {
	ID value.ID
}

// UpdateDietInput is the input for the Update method.
type UpdateDietInput struct {
	Diet diet.Diet
}

// Diet is the interface that wraps the basic Diet methods.
type Diet interface {
	// All returns all diets.
	All(ctx context.Context, input AllDietInput) ([]*diet.Diet, error)

	// Delete deletes the diet with the given id.
	Delete(ctx context.Context, input DeleteDietInput) error

	// Find returns the diet with the given id.
	Find(ctx context.Context, input FindDietInput) (*diet.Diet, error)

	// Save saves the given diet.
	Save(ctx context.Context, input SaveDietInput) (value.ID, error)

	// Update updates the given diet.
	Update(ctx context.Context, input UpdateDietInput) (value.ID, error)
}
