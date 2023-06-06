package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// AllDietsHandler is the handler responsible for getting all diets.
type AllDietsHandler = query.Handler[*AllDietsInput, *queryer.PaginationOutput[*AllDietsOutput]]

// allDietsHandlerImpl is the handler responsible for getting all diets.
type allDietsHandlerImpl struct {
	repo.Diet
}

// NewAllDietsHandler instantiates the handler responsible for getting all diets.
func NewAllDietsHandler(dietRepo repo.Diet) AllDietsHandler {
	errutil.MustBeNotEmpty("repo.Diet", dietRepo)

	return &allDietsHandlerImpl{dietRepo}
}

// Handle implements query.Handler.
func (q *allDietsHandlerImpl) Handle(
	ctx context.Context,
	input *AllDietsInput,
) (*queryer.PaginationOutput[*AllDietsOutput], error) {
	pagination, err := q.Diet.All(ctx, repo.AllDietsInput{
		Filterer:  input.Filter,
		Paginator: &input.Pagination,
		Sorter:    input.Sort,
	})
	if err != nil {
		return nil, err
	}

	output := &queryer.PaginationOutput[*AllDietsOutput]{
		Total:   pagination.Total,
		Results: make([]*FindDietOutput, len(pagination.Results)),
	}

	for i, diet := range pagination.Results {
		output.Results[i] = &FindDietOutput{
			ID:              diet.ID,
			Name:            diet.Name,
			Description:     diet.Description,
			DurationInWeeks: diet.DurationInWeeks,
			Goal:            diet.Goal,
			MealPlan:        diet.MealPlan,
			MonthlyCostUSD:  diet.MonthlyCostUSD,
		}
	}

	return output, nil
}
