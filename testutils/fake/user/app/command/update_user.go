package fake

import (
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func UpdateUserInput() *command.UpdateUserInput {
	data := new(command.UpdateUserInput)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
