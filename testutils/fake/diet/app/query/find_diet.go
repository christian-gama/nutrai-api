package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func FindDietInput() *query.FindDietInput {
	data := new(query.FindDietInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func FindDietOutput() *query.FindDietOutput {
	data := new(query.FindDietOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
