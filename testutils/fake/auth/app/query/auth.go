package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AuthInput() *query.JwtAuthInput {
	data := new(query.JwtAuthInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AuthOutput() *query.JwtAuthOutput {
	data := new(query.JwtAuthOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
