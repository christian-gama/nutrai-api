package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// FindDietHandler is the handler responsible for getting a diet.
type FindDietHandler = query.Handler[*FindDietInput, *FindDietOutput]

// findDietHandlerImpl is the handler responsible for getting a diet.
type findDietHandlerImpl struct {
	repo.Diet
}

// NewFindDietHandler instantiates the handler responsible for getting a diet.
func NewFindDietHandler(dietRepo repo.Diet) FindDietHandler {
	errutil.MustBeNotEmpty("repo.Diet", dietRepo)

	return &findDietHandlerImpl{dietRepo}
}

// Handle implements query.Handler.
func (q *findDietHandlerImpl) Handle(
	ctx context.Context,
	input *FindDietInput,
) (*FindDietOutput, error) {
	diet, err := q.Diet.Find(
		ctx,
		repo.FindDietInput{ID: input.ID, Preloader: input.Preload},
	)
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&FindDietOutput{}, diet), nil
}
