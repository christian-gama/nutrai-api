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
	ID        coreValue.ID   `faker:"uint"`
	WeightKG  value.WeightKG `faker:"boundary_start=1, boundary_end=999"`
	HeightM   value.HeightM  `faker:"boundary_start=1, boundary_end=3"`
	Age       value.Age      `faker:"boundary_start=1, boundary_end=100"`
	BMI       value.BMI      `faker:"boundary_start=16, boundary_end=30"`
	Allergies []*Allergy
}

// NewPatient returns a new patient instance.
func NewPatient() *Patient {
	return &Patient{}
}

// String implements the fmt.Stringer interface.
func (Patient) String() string {
	return "patient"
}

// Validate returns an error if the patient is invalid.
func (p Patient) Validate() (*Patient, error) {
	var errs *errutil.Error

	if p.Age == 0 {
		errs = errutil.Append(errs, errutil.Required("Age"))
	}

	if p.HeightM == 0 {
		errs = errutil.Append(errs, errutil.Required("HeightM"))
	}

	if p.WeightKG == 0 {
		errs = errutil.Append(errs, errutil.Required("WeightKG"))
	}

	for _, allergy := range p.Allergies {
		if _, err := allergy.Validate(); err != nil {
			errs = errutil.Append(errs, err)
		}
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return &p, nil
}

// SetID sets the ID on the builder.
func (p *Patient) SetID(id coreValue.ID) *Patient {
	p.ID = id
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

// SetAllergies sets the allergies on the builder. It will also set the patient ID on each allergy
// and check if the allergy already exists on the patient. If it does not exist, it will create a
// new allergy - otherwise, it will use the existing allergy. It will also remove any allergies
// that are not present in the new list.
func (p *Patient) SetAllergies(allergies []*Allergy) *Patient {
	allergiesMap := make(map[value.Allergy]*Allergy, len(p.Allergies))
	for _, allergy := range p.Allergies {
		allergiesMap[allergy.Name] = allergy
	}

	p.Allergies = make([]*Allergy, len(allergies))
	for i, allergy := range allergies {
		if allergy, ok := allergiesMap[allergy.Name]; ok {
			p.Allergies[i] = allergy
			continue
		}

		p.Allergies[i] = allergy.SetPatientID(p.ID)
	}

	return p
}
