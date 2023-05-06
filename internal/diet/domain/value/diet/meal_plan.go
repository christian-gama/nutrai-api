package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// MealPlan represents the meal plan of a diet.
type MealPlan string

// String returns the string representation of the meal plan.
func (m MealPlan) String() string {
	return string(m)
}

// Validate returns an error if the meal plan is invalid.
func (m MealPlan) Validate() error {
	const fieldName = "MealPlan"

	validMealPlans := []MealPlan{
		Ketogenic,
		Vegetarian,
		Vegan,
		Mediterranean,
		Paleolithic,
		LowCarb,
	}

	if len(m) == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	for _, validMealPlan := range validMealPlans {
		if validMealPlan == m {
			return nil
		}
	}

	return errutil.NewErrInvalid(fieldName, fmt.Sprintf("must be one of %v", validMealPlans))
}

const (
	Ketogenic     MealPlan = "ketogenic"
	Vegetarian    MealPlan = "vegetarian"
	Vegan         MealPlan = "vegan"
	Mediterranean MealPlan = "mediterranean"
	Paleolithic   MealPlan = "paleolithic"
	LowCarb       MealPlan = "low-carb"
)
