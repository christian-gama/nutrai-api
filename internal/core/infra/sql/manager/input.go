package manager

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// SaveInput is the input for Save method.
type SaveInput[Model any] struct {
	Model *Model
}

// FindInput is the input for Find method.
type FindInput[Model any] struct {
	ID value.ID
	queryer.Filterer
	queryer.Preloader
}

// DeleteInput is the input for Delete method.
type DeleteInput[Model any] struct {
	IDs []value.ID
}

// AllInput is the input for All method.
type AllInput[Model any] struct {
	queryer.Filterer
	queryer.Paginator
	queryer.Sorter
	queryer.Preloader
}

// UpdateInput is the input for Update method.
type UpdateInput[Model any] struct {
	Model *Model
	ID    value.ID
}
