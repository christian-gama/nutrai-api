package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func SaveUserInput() *command.SaveUserInput {
	data := new(command.SaveUserInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
