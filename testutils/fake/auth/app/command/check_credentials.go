package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func CheckCredentialsInput() *command.CheckCredentialsInput {
	data := new(command.CheckCredentialsInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
