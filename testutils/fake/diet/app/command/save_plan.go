package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func SavePlanInput() *command.SavePlanInput {
	data := new(command.SavePlanInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
