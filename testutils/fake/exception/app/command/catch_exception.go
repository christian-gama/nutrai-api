package fake

import (
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func CatchExceptionInput() *command.CatchExceptionInput {
	data := new(command.CatchExceptionInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
