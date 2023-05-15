package patient

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
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
	User     *user.User     `faker:"-"`
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

	if p.User == nil {
		errs = errutil.Append(errs, errutil.NewErrRequired("user"))
	} else if err := p.User.Validate(); err != nil {
		errs = errutil.Append(errs, err)
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}

type builder struct {
	patient *Patient
}

// NewBuilder creates a new builder for a patient.
func NewBuilder() *builder {
	return &builder{
		patient: &Patient{},
	}
}

// SetID sets the ID on the builder.
func (b *builder) SetID(id coreValue.ID) *builder {
	b.patient.ID = id
	return b
}

// SetUser sets the user on the builder.
func (b *builder) SetUser(user *user.User) *builder {
	b.patient.User = user
	return b
}

// SetWeightKG sets the weightKG on the builder.
func (b *builder) SetWeightKG(weightKG value.WeightKG) *builder {
	b.patient.WeightKG = weightKG
	return b
}

// SetHeightM sets the heightM on the builder.
func (b *builder) SetHeightM(heightM value.HeightM) *builder {
	b.patient.HeightM = heightM
	return b
}

// SetAge sets the age on the builder.
func (b *builder) SetAge(age value.Age) *builder {
	b.patient.Age = age
	return b
}

// Build builds the patient.
func (b *builder) Build() (*Patient, error) {
	if err := b.patient.Validate(); err != nil {
		return nil, err
	}

	return b.patient, nil
}
