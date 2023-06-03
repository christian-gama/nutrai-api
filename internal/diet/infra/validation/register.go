package validation

import "github.com/christian-gama/nutrai-api/internal/core/infra/validation"

// Register is the function that registers the validation rules for this module.
func Register() {
	validation.RegisterAlias("diet_name", "min=0,max=100")
	validation.RegisterAlias("diet_description", "min=0,max=500")
	validation.RegisterAlias("diet_duration_in_weeks", "number,min=1,max=520")
	validation.RegisterAlias("diet_meal_plan", "min=0,max=100")
	validation.RegisterAlias("diet_monthly_cost_usd", "number,min=0")
	validation.RegisterAlias("diet_goal", "min=0,max=100")
	validation.RegisterAlias("diet_restricted_food_name", "min=0,max=100")
	validation.RegisterAlias("diet_patient_id", "number,min=1")
}
