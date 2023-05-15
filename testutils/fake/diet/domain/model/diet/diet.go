package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	restrictfood "github.com/christian-gama/nutrai-api/internal/diet/domain/model/restrictedfood"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	restrictedFoodFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/restrictedfood"
	"github.com/go-faker/faker/v4"
)

func Diet() *diet.Diet {
	data := new(diet.Diet)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	data.RestrictedFood = []*restrictfood.RestrictedFood{
		restrictedFoodFake.RestrictedFood(),
		restrictedFoodFake.RestrictedFood(),
	}
	data.Goal = value.WeightLoss
	data.MealPlan = value.Vegetarian

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
