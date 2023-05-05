package fake

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/go-faker/faker/v4"
)

func Diet() *diet.Diet {
	data := new(diet.Diet)

	err := faker.FakeData(data)

	if err != nil {
		panic(fmt.Errorf("error while generating fake Diet: %w", err))
	}

	data.AllowedFood = []value.AllowedFood{value.AllowedFood(faker.Name()), value.AllowedFood(faker.Name())}

	data.RestrictedFood = []value.RestrictedFood{value.RestrictedFood(faker.Name()), value.RestrictedFood(faker.Name())}

	if err := data.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake Diet: %w", err))
	}

	return data
}
