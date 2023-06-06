package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AuthJwtInput() *query.AuthJwtInput {
	data := new(query.AuthJwtInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AuthJwtOutput() *query.AuthJwtOutput {
	data := new(query.AuthJwtOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
