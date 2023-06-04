package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func DeletePlanInput() *command.DeletePlanInput {
	data := new(command.DeletePlanInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
