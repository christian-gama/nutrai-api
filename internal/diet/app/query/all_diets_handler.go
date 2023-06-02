package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
)

// AllDietsHandler is the handler responsible for getting all diets.
type AllDietsHandler = query.Handler[*AllDietsInput, *queryer.PaginationOutput[*AllDietsOutput]]

// allDietsHandlerImpl is the handler responsible for getting all diets.
type allDietsHandlerImpl struct {
	repo.Diet
}

// NewAllDietsHandler instantiates the handler responsible for getting all diets.
func NewAllDietsHandler(dietRepo repo.Diet) AllDietsHandler {
	if dietRepo == nil {
		panic(errors.New("repo.Diet cannot be nil"))
	}
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
		Results: []*FindDietOutput{},
	}

	for _, diet := range pagination.Results {
		output.Results = append(output.Results, &FindDietOutput{
			ID:              diet.ID,
			Name:            diet.Name,
			Description:     diet.Description,
			DurationInWeeks: diet.DurationInWeeks,
			Goal:            diet.Goal,
			MealPlan:        diet.MealPlan,
			MonthlyCostUSD:  diet.MonthlyCostUSD,
		})
	}

	return output, nil
}
