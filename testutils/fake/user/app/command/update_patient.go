package fake

import (
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func UpdatePatientInput() *command.UpdatePatientInput {
	data := new(command.UpdatePatientInput)
	data.User = UpdateUserInput()

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
