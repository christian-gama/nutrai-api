package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
)

// SaveDietInput represents the input data for the SaveDiet command.
type SaveDietInput struct {
	User *user.User `ctx:"currentUser" json:"-"`

	Name            value.Name             `json:"name" validate:"required, diet_name" faker:"len=50"`
	Description     value.Description      `json:"description" validate:"required, diet_description" faker:"len=100"`
	DurationInWeeks value.DurationInWeeks  `json:"durationInWeeks" validate:"required, diet_duration_in_weeks" faker:"boundary_start=1,boundary_end=52"`
	Goal            value.Goal             `json:"goal" validate:"required, diet_goal" faker:"-"`
	MealPlan        value.MealPlan         `json:"mealPlan" validate:"required, diet_meal_plan" faker:"-"`
	MonthlyCostUSD  value.MonthlyCostUSD   `json:"monthlyCostUSD" validate:"required, diet_monthly_cost_usd" faker:"boundary_start=12.65, boundary_end=184.05"`
	PatientID       coreValue.ID           `json:"patientId" validate:"required, diet_patient_id" faker:"boundary_start=1,boundary_end=100"`
	RestrictedFood  []value.RestrictedFood `json:"restrictedFood" validate:"required,dive,diet_restricted_food" faker:"-"`
}
