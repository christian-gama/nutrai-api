package manager

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
)

// Save is a generic interface that represents the save action of any model.
type Save[Model any] interface {
	Save(ctx context.Context, input SaveInput[Model]) (*Model, error)
}

// Find is a generic interface that represents the find action of any model. It
// returns the model and an error if it does not exist.
type Find[Model any] interface {
	Find(ctx context.Context, input FindInput[Model]) (*Model, error)
}

// All is a generic interface that represents the all action of any model. It
// returns a pagination output or an error if something goes wrong.
type All[Model any] interface {
	All(ctx context.Context, input AllInput[Model]) (*queryer.PaginationOutput[*Model], error)
}

// Update is a generic interface that represents the update action of any model.
type Update[Model any] interface {
	Update(ctx context.Context, input UpdateInput[Model]) error
}

// Delete is a generic interface that represents the delete action of any model.
type Delete[Model any] interface {
	Delete(ctx context.Context, input DeleteInput[Model]) error
}

// Repository is a collection of interfaces that represents the main actions of
// any model.
type Repository[Model any] interface {
	Save[Model]
	Find[Model]
	All[Model]
	Update[Model]
	Delete[Model]
}
