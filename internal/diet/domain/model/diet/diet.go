package diet

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Diet represents a Diet model, which includes information about a specific diet plan.
// It includes attributes such as the diet plan's name, description, duration, goals, allowed and
// restricted foods, meal plan, nutritional information and cost.
// This model can be used to represent any type of diet plan, such as a low-carb diet, vegan diet,
// or Mediterranean diet.
type Diet struct {
	ID              coreValue.ID          `faker:"uint"`
	PatientID       coreValue.ID          `faker:"uint"`
	Name            value.Name            `faker:"name"`
	Description     value.Description     `faker:"sentence"`
	DurationInWeeks value.DurationInWeeks `faker:"boundary_start=1, boundary_end=100"`
	Goal            value.Goal            `faker:"-"`
	MealPlan        value.MealPlan        `faker:"-"`
	MonthlyCostUSD  value.MonthlyCostUSD  `faker:"boundary_start=12.65, boundary_end=184.05"`
}

// NewDiet returns a new Diet instance.
func NewDiet() *Diet {
	return &Diet{}
}

// Validate returns an error if the diet is invalid.
func (d *Diet) Validate() (*Diet, error) {
	var errs *errutil.Error

	if err := d.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := d.PatientID.Validate(); err != nil {
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

	if errs.HasErrors() {
		return nil, errs
	}

	return d, nil
}

// SetID sets the diet's ID.
func (d *Diet) SetID(id coreValue.ID) *Diet {
	d.ID = id
	return d
}

// SetPatientID sets the diet's patient ID.
func (d *Diet) SetPatientID(patientID coreValue.ID) *Diet {
	d.PatientID = patientID
	return d
}

// SetName sets the diet's name.
func (d *Diet) SetName(name value.Name) *Diet {
	d.Name = name
	return d
}

// SetDescription sets the diet's description.
func (d *Diet) SetDescription(description value.Description) *Diet {
	d.Description = description
	return d
}

// SetDurationInWeeks sets the diet's duration in weeks.
func (d *Diet) SetDurationInWeeks(durationInWeeks value.DurationInWeeks) *Diet {
	d.DurationInWeeks = durationInWeeks
	return d
}

// SetGoal sets the diet's goal.
func (d *Diet) SetGoal(goal value.Goal) *Diet {
	d.Goal = goal
	return d
}

// SetMealPlan sets the diet's meal plan.
func (d *Diet) SetMealPlan(mealPlan value.MealPlan) *Diet {
	d.MealPlan = mealPlan
	return d
}

// SetMonthlyCostUSD sets the diet's monthly cost in USD.
func (d *Diet) SetMonthlyCostUSD(monthlyCostUSD value.MonthlyCostUSD) *Diet {
	d.MonthlyCostUSD = monthlyCostUSD
	return d
}
