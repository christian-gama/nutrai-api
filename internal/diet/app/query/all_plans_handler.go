package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

// AllPlansHandler is the handler responsible for getting all plans.
type AllPlansHandler = query.Handler[*AllPlansInput, *queryer.PaginationOutput[*AllPlansOutput]]

// allPlansHandlerImpl is the handler responsible for getting all plans.
type allPlansHandlerImpl struct {
	repo.Plan
}

// NewAllPlansHandler instantiates the handler responsible for getting all plans.
func NewAllPlansHandler(planRepo repo.Plan) AllPlansHandler {
	if planRepo == nil {
		panic(errors.New("repo.Plan cannot be nil"))
	}
	return &allPlansHandlerImpl{planRepo}
}

// Handle implements query.Handler.
func (q *allPlansHandlerImpl) Handle(
	ctx context.Context,
	input *AllPlansInput,
) (*queryer.PaginationOutput[*AllPlansOutput], error) {
	pagination, err := q.Plan.All(ctx, repo.AllPlansInput{
		Filterer:  input.Filter,
		Paginator: &input.Pagination,
		Sorter:    input.Sort,
		Preloader: input.Preload,
	})
	if err != nil {
		return nil, err
	}

	output := &queryer.PaginationOutput[*AllPlansOutput]{
		Total:   pagination.Total,
		Results: []*FindPlanOutput{},
	}

	for _, plan := range pagination.Results {
		var diet *FindDietOutput
		if plan.Diet != nil {
			diet = &FindDietOutput{
				ID:              plan.Diet.ID,
				Name:            plan.Diet.Name,
				Description:     plan.Diet.Description,
				DurationInWeeks: plan.Diet.DurationInWeeks,
				Goal:            plan.Diet.Goal,
				MealPlan:        plan.Diet.MealPlan,
				MonthlyCostUSD:  plan.Diet.MonthlyCostUSD,
			}
		}

		output.Results = append(output.Results, &FindPlanOutput{
			ID:     plan.ID,
			DietID: plan.DietID,
			Text:   plan.Text,
			Diet:   diet,
		})
	}

	return output, nil
}
