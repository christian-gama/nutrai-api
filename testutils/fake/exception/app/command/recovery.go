package fake

import (
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func RecoveryInput() *command.RecoveryInput {
	data := new(command.RecoveryInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
