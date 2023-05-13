package manager

import (
	queryingPort "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
)

// SaveInput is the input for Save method.
type SaveInput[Model any] struct {
	Model *Model
}

// FindInput is the input for Find method.
type FindInput[Model any] struct {
	ID value.ID
	queryingPort.Filterer
	queryingPort.Preloader
}

// DeleteInput is the input for Delete method.
type DeleteInput[Model any] struct {
	IDs []value.ID
}

// AllInput is the input for All method.
type AllInput[Model any] struct {
	queryingPort.Filterer
	queryingPort.Paginator
	queryingPort.Sorter
	queryingPort.Preloader
}

// UpdateInput is the input for Update method.
type UpdateInput[Model any] struct {
	Model *Model
	ID    value.ID
}
