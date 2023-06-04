package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func LoginInput() *service.LoginInput {
	data := new(service.LoginInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func LoginOutput() *service.LoginOutput {
	data := new(service.LoginOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
