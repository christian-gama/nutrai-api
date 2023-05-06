package fake

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/go-faker/faker/v4"
)

func Patient() *patient.Patient {
	data := new(patient.Patient)

	err := faker.FakeData(data)
	if err != nil {
		panic(fmt.Errorf("error while generating fake Patient: %w", err))
	}

	if err := data.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake Patient: %w", err))
	}

	return data
}
