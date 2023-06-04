package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// SavePlanHandler represents the SavePlan command.
type SavePlanHandler = service.Handler[*SavePlanInput, *SavePlanOutput]

// savePlanHandlerImpl represents the implementation of the SavePlan command.
type savePlanHandlerImpl struct {
	repo.Plan
	repo.Diet
}

// NewSavePlanHandler returns a new Save instance.
func NewSavePlanHandler(
	planRepo repo.Plan,
	dietRepo repo.Diet,
) SavePlanHandler {
	errutil.MustBeNotEmpty("repo.Plan", planRepo)
	errutil.MustBeNotEmpty("repo.Diet", dietRepo)

	return &savePlanHandlerImpl{planRepo, dietRepo}
}

// Handle implements command.Handler.
func (c *savePlanHandlerImpl) Handle(
	ctx context.Context,
	input *SavePlanInput,
) (*SavePlanOutput, error) {
	savedDiet, err := c.Diet.Find(
		ctx,
		repo.FindDietInput{
			ID:       input.DietID,
			Filterer: querying.AddFilter("patientID", querying.EqOperator, input.User.ID),
		},
	)
	if err != nil {
		return nil, err
	}

	p, err := plan.NewPlan().
		SetDietID(savedDiet.ID).
		SetText("replace with AI Generated data").
		Validate()
	if err != nil {
		return nil, err
	}

	p, err = c.Plan.Save(ctx, repo.SavePlanInput{Plan: p})
	if err != nil {
		return nil, err
	}

	return &SavePlanOutput{
		ID:     p.ID,
		DietID: p.DietID,
		Text:   p.Text,
	}, nil
}
