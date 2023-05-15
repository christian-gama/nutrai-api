package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
)

// SaveDietInput is the input for the Save method.
type SaveDietInput struct {
	Diet *diet.Diet
}

// AllDietsInput is the input for the All method.
type AllDietsInput struct {
	queryer.Filterer
	queryer.Sorter
	queryer.Paginator
	queryer.Preloader
}

// FindDietInput is the input for the Find method.
type FindDietInput struct {
	ID value.ID
	queryer.Filterer
	queryer.Preloader
}

// DeleteDietInput is the input for the Delete method.
type DeleteDietInput struct {
	IDs []value.ID
}

// UpdateDietInput is the input for the Update method.
type UpdateDietInput struct {
	Diet *diet.Diet
	ID   value.ID
}

// Diet is the interface that wraps the basic Diet methods.
type Diet interface {
	// All returns all diets.
	All(ctx context.Context, input AllDietsInput) (*queryer.PaginationOutput[*diet.Diet], error)

	// Delete deletes the diet with the given id.
	Delete(ctx context.Context, input DeleteDietInput) error

	// Find returns the diet with the given id.
	Find(ctx context.Context, input FindDietInput) (*diet.Diet, error)

	// Save saves the given diet.
	Save(ctx context.Context, input SaveDietInput) (*diet.Diet, error)

	// Update updates the given diet.
	Update(ctx context.Context, input UpdateDietInput) error
}
