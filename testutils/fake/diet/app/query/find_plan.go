package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func FindPlanInput() *query.FindPlanInput {
	data := new(query.FindPlanInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func FindPlanOutput() *query.FindPlanOutput {
	data := new(query.FindPlanOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
