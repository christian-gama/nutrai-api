package plan

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Plan represents a detailed blueprint of a specific diet regimen, linked to the Diet model through
// the DietID. This model provides a comprehensive plan that helps individuals to effectively follow
// the diet.
type Plan struct {
	ID     coreValue.ID `faker:"uint"`
	DietID coreValue.ID `faker:"uint"`
	Diet   *diet.Diet   `faker:"-"`
	Text   value.Plan   `faker:"paragraph"`
}

// NewPlan creates a new Plan instance.
func NewPlan() *Plan {
	return &Plan{}
}

// String implements the fmt.Stringer interface.
func (Plan) String() string {
	return "plan"
}

// Validate returns an error if the plan is invalid.
func (p *Plan) Validate() (*Plan, error) {
	var errs *errutil.Error

	if err := p.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.DietID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.Text.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if p.Diet == nil {
		errs = errutil.Append(errs, errutil.Required("diet"))
	} else if _, err := p.Diet.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs != nil {
		return nil, errs
	}

	return p, nil
}

// SetID sets the ID of the Plan.
func (p *Plan) SetID(id coreValue.ID) *Plan {
	p.ID = id
	return p
}

// SetDietID sets the DietID of the Plan.
func (p *Plan) SetDietID(dietID coreValue.ID) *Plan {
	p.DietID = dietID
	return p
}

// SetDiet sets the Diet of the Plan.
func (p *Plan) SetDiet(diet *diet.Diet) *Plan {
	p.Diet = diet
	return p
}

// SetText sets the Text of the Plan.
func (p *Plan) SetText(text value.Plan) *Plan {
	p.Text = text
	return p
}
