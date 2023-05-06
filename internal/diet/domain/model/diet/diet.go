package diet

import (
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Diet represents a Diet model, which includes information about a specific diet plan.
// It includes attributes such as the diet plan's name, description, duration, goals, allowed and
// restricted foods, meal plan, nutritional information and cost.
// This model can be used to represent any type of diet plan, such as a low-carb diet, vegan diet,
// or Mediterranean diet.
type Diet struct {
	ID              sharedvalue.ID         `faker:"uint"`
	Name            value.Name             `faker:"name"`
	Description     value.Description      `faker:"paragraph"`
	RestrictedFood  []value.RestrictedFood `faker:"-"`
	DurationInWeeks value.DurationInWeeks  `faker:"boundary_start=1, boundary_end=100"`
	Goal            value.Goal             `faker:"oneof:WEIGHT_LOSS, WEIGHT_GAIN, MAINTAIN, IMPROVED_HEALTH, INCREASE_ENERGY, MUSCLE_GAIN"`
	MealPlan        value.MealPlan         `faker:"oneof:ketogenic, vegetarian, vegan, mediterranean, paleolithic, low-carb"`
	MonthlyCostUSD  value.MonthlyCostUSD   `faker:"boundary_start=12.65, boundary_end=184.05"`
}

// DietInput represents the input to create a new Diet.
func NewDiet(input *DietInput) (*Diet, error) {
	diet := Diet(*input)

	if err := diet.Validate(); err != nil {
		return nil, err
	}

	return &diet, nil
}

// Validate returns an error if the diet is invalid.
func (d Diet) Validate() error {
	var errs *errutil.Error

	if err := d.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.Name.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.Description.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.DurationInWeeks.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.Goal.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.MealPlan.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.MonthlyCostUSD.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if len(d.RestrictedFood) == 0 {
		errs = errutil.Append(errs, errutil.NewErrRequired("restricted_food"))
	}

	for _, restrictedFood := range d.RestrictedFood {
		if err := restrictedFood.Validate(); err != nil {
			errs = errutil.Append(errs, err)
		}
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
