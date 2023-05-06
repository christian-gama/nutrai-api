package diet

import (
	"errors"

	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
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

func NewDiet(
	input DietInput,
) (*Diet, error) {
	diet := Diet(input)

	if err := diet.Validate(); err != nil {
		return nil, err
	}

	return &diet, nil
}

func (d Diet) Validate() error {
	if isValid := d.ID.IsValid(); !isValid {
		return errors.New("invalid ID")
	}

	if isValid := d.Name.IsValid(); !isValid {
		return errors.New("invalid name")
	}

	if isValid := d.Description.IsValid(); !isValid {
		return errors.New("invalid description")
	}

	if isValid := d.DurationInWeeks.IsValid(); !isValid {
		return errors.New("invalid duration in weeks")
	}

	if isValid := d.Goal.IsValid(); !isValid {
		return errors.New("invalid goal")
	}

	if isValid := d.MealPlan.IsValid(); !isValid {
		return errors.New("invalid meal plan")
	}

	if isValid := d.MonthlyCostUSD.IsValid(); !isValid {
		return errors.New("invalid monthly cost")
	}

	for _, restrictedFood := range d.RestrictedFood {
		if isValid := restrictedFood.IsValid(); !isValid {
			return errors.New("invalid restricted food")
		}
	}

	return nil
}
