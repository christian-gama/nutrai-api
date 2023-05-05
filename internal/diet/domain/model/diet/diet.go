package diet

import (
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
)

// Diet represents a Diet model, which includes information about a specific diet plan.
// It includes attributes such as the diet plan's name, description, duration, goals, allowed and
// restricted foods, meal plan, nutritional information and cost.
// This model can be used to represent any type of diet plan, such as a low-carb diet, vegan diet,
// or Mediterranean diet.
type Diet struct {
	ID              sharedvalue.ID
	Name            value.Name
	Description     value.Description
	AllowedFood     []value.AllowedFood
	RestrictedFood  []value.RestrictedFood
	DurationInWeeks value.DurationInWeeks
	Goal            value.Goal
	MealPlan        value.MealPlan
	MonthlyCostUSD  value.MonthlyCostUSD
}
