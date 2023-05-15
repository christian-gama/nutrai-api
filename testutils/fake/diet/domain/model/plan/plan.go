package plan

import (
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Plan() *plan.Plan {
	data := new(plan.Plan)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
