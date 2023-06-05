package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AuthInput() *query.AuthInput {
	data := new(query.AuthInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AuthOutput() *query.AuthOutput {
	data := new(query.AuthOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
