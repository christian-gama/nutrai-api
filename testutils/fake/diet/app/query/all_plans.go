package fake

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	"github.com/go-faker/faker/v4"
)

func AllPlansInput() *query.AllPlansInput {
	data := new(query.AllPlansInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AllPlansOutput() *queryer.PaginationOutput[*plan.Plan] {
	data := new(queryer.PaginationOutput[*plan.Plan])

	data.Results = []*plan.Plan{fake.Plan()}
	data.Total = 1

	return data
}
