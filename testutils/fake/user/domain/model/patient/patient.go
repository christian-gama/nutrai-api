package fake

import (
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Patient() *patient.Patient {
	data := new(patient.Patient)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating("patient", err)
	}

	if err := data.Validate(); err != nil {
		fake.ErrGenerating("patient", err)
	}

	return data
}
