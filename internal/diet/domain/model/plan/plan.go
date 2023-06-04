package plan

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/plan"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Plan represents a detailed blueprint of a specific diet regimen, linked to the Diet model through
// the DietID. This model provides a comprehensive plan that helps individuals to effectively follow
// the diet.
type Plan struct {
	ID     coreValue.ID `faker:"uint"`
	DietID coreValue.ID `faker:"uint"`
	Text   value.Text   `faker:"paragraph"`
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

	if err := p.DietID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if p.Text == "" {
		errs = errutil.Append(errs, errors.Required("text"))
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

// SetText sets the Text of the Plan.
func (p *Plan) SetText(text value.Text) *Plan {
	p.Text = text
	return p
}
