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

// New creates a new restricted food model.
func New() *RestrictedFood {
	return &RestrictedFood{}
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

// Build builds and returns the restricted food model.
func (rf *RestrictedFood) Build() (*RestrictedFood, error) {
	if err := rf.Validate(); err != nil {
		return nil, err
	}

	return rf, nil
}
