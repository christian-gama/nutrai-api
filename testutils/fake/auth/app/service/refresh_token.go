package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func RefreshTokenInput() *service.RefreshTokenInput {
	data := new(service.RefreshTokenInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func RefreshTokenOutput() *service.RefreshTokenOutput {
	data := new(service.RefreshTokenOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
