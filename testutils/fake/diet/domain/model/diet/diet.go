package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	patientFake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"

	"github.com/go-faker/faker/v4"
)

func Diet() *diet.Diet {
	data := new(diet.Diet)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	data.Patient = patientFake.Patient()
	data.RestrictedFood = []value.RestrictedFood{value.RestrictedFood(faker.Name()), value.RestrictedFood(faker.Name())}
	data.Goal = value.WeightLoss
	data.MealPlan = value.Vegetarian

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
