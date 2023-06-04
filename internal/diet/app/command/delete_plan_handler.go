package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// DeletePlanHandler represents the DeletePlan command.
type DeletePlanHandler = command.Handler[*DeletePlanInput]

// deletePlanHandlerImpl represents the implementation of the DeletePlan command.
type deletePlanHandlerImpl struct {
	repo.Plan
}

// NewDeletePlanHandler returns a new Delete instance.
func NewDeletePlanHandler(repoPlan repo.Plan) DeletePlanHandler {
	errutil.MustBeNotEmpty("repo.Plan", repoPlan)

	return &deletePlanHandlerImpl{repoPlan}
}

// Handle implements command.Handler.
func (c *deletePlanHandlerImpl) Handle(ctx context.Context, input *DeletePlanInput) error {
	_, err := c.Find(ctx, repo.FindPlanInput{ID: input.User.ID})
	if err != nil {
		return err
	}

	return c.Delete(ctx, repo.DeletePlanInput{IDs: []value.ID{input.User.ID}})
}
