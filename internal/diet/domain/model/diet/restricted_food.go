package diet

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// RestrictedFood represents a model that details the foods which are restricted or not allowed in a
// specific diet plan. This model is related to the Diet model through the DietID. Each Diet can
// have multiple RestrictedFood entries, allowing for the representation of complex dietary
// restrictions.
// The Name attribute specifies the name of the food that is restricted in the diet plan.
type RestrictedFood struct {
	ID     coreValue.ID         `faker:"uint"`
	DietID coreValue.ID         `faker:"uint"`
	Name   value.RestrictedFood `faker:"-"`
}

// New creates a new RestrictedFood instance.
func NewRestrictedFood() *RestrictedFood {
	return &RestrictedFood{}
}

// String implements the fmt.Stringer interface.
func (RestrictedFood) String() string {
	return "restricted food"
}

// Validate returns an error if the restricted food is invalid.
func (rf *RestrictedFood) Validate() (*RestrictedFood, error) {
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
		return nil, errs
	}

	return rf, nil
}

// SetID sets the ID field for the restricted food model.
func (rf *RestrictedFood) SetID(id coreValue.ID) *RestrictedFood {
	rf.ID = id
	return rf
}

// SetDietID sets the DietID field for the restricted food model.
func (rf *RestrictedFood) SetDietID(dietID coreValue.ID) *RestrictedFood {
	rf.DietID = dietID
	return rf
}

// SetName sets the Name field for the restricted food model.
func (rf *RestrictedFood) SetName(name value.RestrictedFood) *RestrictedFood {
	rf.Name = name
	return rf
}
