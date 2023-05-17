package diet

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type RestrictedFood struct {
	ID     coreValue.ID         `faker:"uint"`
	DietID coreValue.ID         `faker:"uint"`
	Name   value.RestrictedFood `faker:"-"`
}

// Validate returns an error if the restricted food is invalid.
func (rf *RestrictedFood) Validate() error {
	var errs *errutil.Error

	if err := rf.ID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := rf.DietID.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := rf.Name.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}

type builder struct {
	restrictedFood *RestrictedFood
}

// NewBuilder creates a new builder for the restricted food model.
func NewBuilder() *builder {
	return &builder{
		restrictedFood: &RestrictedFood{},
	}
}

// SetID sets the ID field for the restricted food model.
func (b *builder) SetID(id coreValue.ID) *builder {
	b.restrictedFood.ID = id
	return b
}

// SetDietID sets the DietID field for the restricted food model.
func (b *builder) SetDietID(dietID coreValue.ID) *builder {
	b.restrictedFood.DietID = dietID
	return b
}

// SetName sets the Name field for the restricted food model.
func (b *builder) SetName(name value.RestrictedFood) *builder {
	b.restrictedFood.Name = name
	return b
}

// Build builds and returns the restricted food model.
func (b *builder) Build() (*RestrictedFood, error) {
	if err := b.restrictedFood.Validate(); err != nil {
		return nil, err
	}

	return b.restrictedFood, nil
}
