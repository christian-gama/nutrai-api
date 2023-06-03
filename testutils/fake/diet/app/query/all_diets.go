package fake

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	"github.com/go-faker/faker/v4"
)

func AllDietsInput() *query.AllDietsInput {
	data := new(query.AllDietsInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AllDietsOutput() *queryer.PaginationOutput[*diet.Diet] {
	data := new(queryer.PaginationOutput[*diet.Diet])

	data.Results = []*diet.Diet{fake.Diet()}
	data.Total = 1

	return data
}
