package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func SavePatientInput() *command.SavePatientInput {
	data := new(command.SavePatientInput)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
