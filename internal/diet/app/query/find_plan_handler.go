package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// FindPlanHandler is the handler responsible for getting a plan.
type FindPlanHandler = query.Handler[*FindPlanInput, *FindPlanOutput]

// findPlanHandlerImpl is the handler responsible for getting a plan.
type findPlanHandlerImpl struct {
	repo.Plan
}

// NewFindPlanHandler instantiates the handler responsible for getting a plan.
func NewFindPlanHandler(planRepo repo.Plan) FindPlanHandler {
	errutil.MustBeNotEmpty("repo.Plan", planRepo)

	return &findPlanHandlerImpl{planRepo}
}

// Handle implements query.Handler.
func (q *findPlanHandlerImpl) Handle(
	ctx context.Context,
	input *FindPlanInput,
) (*FindPlanOutput, error) {
	plan, err := q.Plan.Find(
		ctx,
		repo.FindPlanInput{ID: input.ID},
	)
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&FindPlanOutput{}, plan), nil
}
