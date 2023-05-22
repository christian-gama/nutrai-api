package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func RestrictedFood() *diet.RestrictedFood {
	data := new(diet.RestrictedFood)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	data.Name = value.RestrictedFood(faker.Name())

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
