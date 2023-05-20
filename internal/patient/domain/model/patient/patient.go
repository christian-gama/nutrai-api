package patient

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Patient represents a Patient model, encapsulating all pertinent information about a patient.
// This includes unique identification, user profile, weight in kilograms, height in meters, age,
// and body mass index (BMI). This model serves as a comprehensive overview of a patient's
// physical characteristics and overall health profile. The user profile could contain additional
// information such as name, contact details, and health history. The weight, height, and BMI
// could be used to calculate dietary needs, track health progress, or establish fitness goals.
type Patient struct {
	ID       coreValue.ID   `faker:"uint"`
	UserID   coreValue.ID   `faker:"uint"`
	WeightKG value.WeightKG `faker:"boundary_start=1, boundary_end=999"`
	HeightM  value.HeightM  `faker:"boundary_start=1, boundary_end=3"`
	Age      value.Age      `faker:"boundary_start=1, boundary_end=100"`
	BMI      value.BMI      `faker:"boundary_start=16, boundary_end=30"`
}

// Validate returns an error if the patient is invalid.
func (p *Patient) Validate() error {
	var errs *errutil.Error

	if err := p.WeightKG.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.HeightM.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if err := p.Age.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}

// New returns a new patient builder.
func New() *Patient {
	return &Patient{}
}

// SetID sets the ID on the builder.
func (p *Patient) SetID(id coreValue.ID) *Patient {
	p.ID = id
	return p
}

// SetUserID sets the user on the builder.
func (p *Patient) SetUserID(userID coreValue.ID) *Patient {
	p.UserID = userID
	return p
}

// SetWeightKG sets the weightKG on the builder.
func (p *Patient) SetWeightKG(weightKG value.WeightKG) *Patient {
	p.WeightKG = weightKG
	return p
}

// SetHeightM sets the heightM on the builder.
func (p *Patient) SetHeightM(heightM value.HeightM) *Patient {
	p.HeightM = heightM
	return p
}

// SetAge sets the age on the builder.
func (p *Patient) SetAge(age value.Age) *Patient {
	p.Age = age
	return p
}

// Build builds the patient.
func (p *Patient) Build() (*Patient, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}
