package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/service"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func SavePlanInput() *service.SavePlanInput {
	data := new(service.SavePlanInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func SavePlanOutput() *service.SavePlanOutput {
	data := new(service.SavePlanOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
