package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Patient() *patient.Patient {
	data := new(patient.Patient)
	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	data.SetAllergies(
		Allergy().SetPatientID(data.ID),
		Allergy().SetPatientID(data.ID),
	)

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
