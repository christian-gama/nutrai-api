package diet

import (
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
)

type DietInput struct {
	ID              sharedvalue.ID
	Name            value.Name
	Description     value.Description
	RestrictedFood  []value.RestrictedFood
	DurationInWeeks value.DurationInWeeks
	Goal            value.Goal
	MealPlan        value.MealPlan
	MonthlyCostUSD  value.MonthlyCostUSD
}
