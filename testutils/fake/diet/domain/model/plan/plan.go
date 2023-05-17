package fake

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	dietFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	"github.com/go-faker/faker/v4"
)

func Plan() *plan.Plan {
	data := new(plan.Plan)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	data.Diet = dietFake.Diet()

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
