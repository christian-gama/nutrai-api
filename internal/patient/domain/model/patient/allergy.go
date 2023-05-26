package patient

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

type Allergy struct {
	ID        coreValue.ID  `faker:"uint"`
	PatientID coreValue.ID  `faker:"uint"`
	Name      value.Allergy `faker:"len=50"`
}

// NewAllergy creates a new Allergy instance.
func NewAllergy() *Allergy {
	return &Allergy{}
}

// String implements the fmt.Stringer interface.
func (Allergy) String() string {
	return "allergy"
}

// Validate validates the Allergy fields.
func (a Allergy) Validate() (*Allergy, error) {
	var errs *errutil.Error

	if a.Name == "" {
		errs = errutil.Append(errs, errors.Required("Name"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return &a, nil
}

// SetID sets the ID field.
func (a *Allergy) SetID(id coreValue.ID) *Allergy {
	a.ID = id
	return a
}

// SetPatientID sets the PatientID field.
func (a *Allergy) SetPatientID(patientID coreValue.ID) *Allergy {
	a.PatientID = patientID
	return a
}

// SetName sets the Name field.
func (a *Allergy) SetName(name value.Allergy) *Allergy {
	a.Name = name
	return a
}
