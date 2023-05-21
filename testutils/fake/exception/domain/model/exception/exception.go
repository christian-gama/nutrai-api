package fake

import (
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Exception() *exception.Exception {
	data := new(exception.Exception)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
