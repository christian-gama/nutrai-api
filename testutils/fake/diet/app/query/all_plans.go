package fake

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
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

func AllPlansOutput() *queryer.PaginationOutput[*query.AllPlansOutput] {
	data := new(queryer.PaginationOutput[*query.AllPlansOutput])

	data.Results = []*query.AllPlansOutput{FindPlanOutput()}
	data.Total = 1

	return data
}
