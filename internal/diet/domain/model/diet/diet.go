package diet

import (
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Diet represents a Diet model, which includes information about a specific diet plan.
// It includes attributes such as the diet plan's name, description, duration, goals, allowed and
// restricted foods, meal plan, nutritional information and cost.
// This model can be used to represent any type of diet plan, such as a low-carb diet, vegan diet,
// or Mediterranean diet.
type Diet struct {
	ID              sharedvalue.ID         `faker:"uint"`
	Patient         *patient.Patient       `faker:"-"`
	Name            value.Name             `faker:"name"`
	Description     value.Description      `faker:"paragraph"`
	RestrictedFood  []value.RestrictedFood `faker:"-"`
	DurationInWeeks value.DurationInWeeks  `faker:"boundary_start=1, boundary_end=100"`
	Goal            value.Goal             `faker:"-"`
	MealPlan        value.MealPlan         `faker:"-"`
	MonthlyCostUSD  value.MonthlyCostUSD   `faker:"boundary_start=12.65, boundary_end=184.05"`
}

// Validate returns an error if the diet is invalid.
func (d *Diet) Validate() error {
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

	if d.Patient == nil {
		errs = errutil.Append(errs, errutil.NewErrRequired("patient"))
	} else if err := d.Patient.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}

type builder struct {
	diet *Diet
}

// NewBuilder creates a new builder for the Diet model.
func NewBuilder() *builder {
	return &builder{
		diet: &Diet{},
	}
}

// SetID sets the diet's ID.
func (b *builder) SetID(id sharedvalue.ID) *builder {
	b.diet.ID = id
	return b
}

// SetPatient sets the diet's patient.
func (b *builder) SetPatient(patient *patient.Patient) *builder {
	b.diet.Patient = patient
	return b
}

// SetName sets the diet's name.
func (b *builder) SetName(name value.Name) *builder {
	b.diet.Name = name
	return b
}

// SetDescription sets the diet's description.
func (b *builder) SetDescription(description value.Description) *builder {
	b.diet.Description = description
	return b
}

// SetDurationInWeeks sets the diet's duration in weeks.
func (b *builder) SetDurationInWeeks(durationInWeeks value.DurationInWeeks) *builder {
	b.diet.DurationInWeeks = durationInWeeks
	return b
}

// SetGoal sets the diet's goal.
func (b *builder) SetGoal(goal value.Goal) *builder {
	b.diet.Goal = goal
	return b
}

// SetMealPlan sets the diet's meal plan.
func (b *builder) SetMealPlan(mealPlan value.MealPlan) *builder {
	b.diet.MealPlan = mealPlan
	return b
}

// SetMonthlyCostUSD sets the diet's monthly cost in USD.
func (b *builder) SetMonthlyCostUSD(monthlyCostUSD value.MonthlyCostUSD) *builder {
	b.diet.MonthlyCostUSD = monthlyCostUSD
	return b
}

// SetRestrictedFood sets the diet's restricted food.
func (b *builder) SetRestrictedFood(restrictedFood []value.RestrictedFood) *builder {
	b.diet.RestrictedFood = restrictedFood
	return b
}

// Build builds and returns the diet.
func (b *builder) Build() (*Diet, error) {
	if err := b.diet.Validate(); err != nil {
		return nil, err
	}

	return b.diet, nil
}
