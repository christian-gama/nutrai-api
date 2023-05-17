package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Diet() *diet.Diet {
	data := new(diet.Diet)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	data.Goal = value.WeightLoss
	data.MealPlan = value.Vegetarian

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
