package patient

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
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
	// There are no fields to validate, but this method is required to keep the
	// consistency with the other models.
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
