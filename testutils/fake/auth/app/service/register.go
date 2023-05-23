package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func RegisterInput() *service.RegisterInput {
	data := new(service.RegisterInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
