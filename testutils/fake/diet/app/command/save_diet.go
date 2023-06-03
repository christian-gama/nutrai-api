package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func SaveDietInput() *command.SaveDietInput {
	data := new(command.SaveDietInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	data.RestrictedFood = []value.RestrictedFood{
		value.RestrictedFood(faker.Name()),
		value.RestrictedFood(faker.Name()),
	}
	data.Goal = value.WeightLoss
	data.MealPlan = value.Vegetarian

	return data
}
