package fake

import (
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Exception() *exception.Exception {
	data := new(exception.Exception)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
