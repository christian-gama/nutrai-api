package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func FindByEmailInput() *query.FindByEmailInput {
	data := new(query.FindByEmailInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func FindByEmailOutput() *query.FindByEmailOutput {
	data := new(query.FindByEmailOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
