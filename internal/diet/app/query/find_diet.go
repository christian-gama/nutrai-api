package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
)

// FindDietInput is the input data of the diet query.
type FindDietInput struct {
	ID coreValue.ID `uri:"id" faker:"uint" validate:"required,min=1"`

	querying.Preload `form:"preload" validate:"query"`
}

// FindDietOutput is the output data of the diet query.
type FindDietOutput struct {
	ID              coreValue.ID          `json:"id"`
	Name            value.Name            `json:"name"`
	Description     value.Description     `json:"description"`
	DurationInWeeks value.DurationInWeeks `json:"durationInWeeks"`
	Goal            value.Goal            `json:"goal"`
	MealPlan        value.MealPlan        `json:"mealPlan"`
	MonthlyCostUSD  value.MonthlyCostUSD  `json:"monthlyCostUSD"`
}
