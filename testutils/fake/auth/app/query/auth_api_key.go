package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AuthApiKeyInput() *query.AuthApiKeyInput {
	data := new(query.AuthApiKeyInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AuthApiKeyOutput() *query.AuthApiKeyOutput {
	data := new(query.AuthApiKeyOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
