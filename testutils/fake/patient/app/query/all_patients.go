package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AllPatientsInput() *query.AllPatientsInput {
	data := new(query.AllPatientsInput)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
