package manager

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
)

type Save[Model any] interface {
	Save(ctx context.Context, input SaveInput[Model]) (*Model, error)
}

type Find[Model any] interface {
	Find(ctx context.Context, input FindInput[Model]) (*Model, error)
}

type All[Model any] interface {
	All(ctx context.Context, input AllInput[Model]) (*queryer.PaginationOutput[*Model], error)
}

type Update[Model any] interface {
	Update(ctx context.Context, input UpdateInput[Model]) error
}

type Delete[Model any] interface {
	Delete(ctx context.Context, input DeleteInput[Model]) error
}

type Repository[Model any] interface {
	Save[Model]
	Find[Model]
	All[Model]
	Update[Model]
	Delete[Model]
}
