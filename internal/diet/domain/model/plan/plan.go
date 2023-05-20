package plan

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Plan represents a Plan model, which includes information about a specific diet plan.
type Plan struct {
	ID     coreValue.ID `faker:"uint"`
	DietID coreValue.ID `faker:"uint"`
	Diet   *diet.Diet   `faker:"-"`
	Text   value.Plan   `faker:"paragraph"`
}

// New creates a new Plan.
func New() *Plan {
	return &Plan{}
}

// Validate returns an error if the plan is invalid.
func (p *Plan) Validate() error {
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
		errs = errutil.Append(errs, errutil.NewErrRequired("diet"))
	} else if err := p.Diet.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs != nil {
		return errs
	}

	return nil
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

// Build returns the built Plan.
func (p *Plan) Build() (*Plan, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}
