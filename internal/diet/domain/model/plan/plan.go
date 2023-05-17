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

type builder struct {
	plan *Plan
}

// NewBuilder creates a new builder for the Plan model.
func NewBuilder() *builder {
	return &builder{
		plan: &Plan{},
	}
}

// SetID sets the ID of the Plan.
func (b *builder) SetID(id coreValue.ID) *builder {
	b.plan.ID = id
	return b
}

// SetDietID sets the DietID of the Plan.
func (b *builder) SetDietID(dietID coreValue.ID) *builder {
	b.plan.DietID = dietID
	return b
}

// SetDiet sets the Diet of the Plan.
func (b *builder) SetDiet(diet *diet.Diet) *builder {
	b.plan.Diet = diet
	return b
}

// SetText sets the Text of the Plan.
func (b *builder) SetText(text value.Plan) *builder {
	b.plan.Text = text
	return b
}

// Build returns the built Plan.
func (b *builder) Build() (*Plan, error) {
	if err := b.plan.Validate(); err != nil {
		return nil, err
	}

	return b.plan, nil
}
