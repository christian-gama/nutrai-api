package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func AllPatientsInput() *query.AllPatientsInput {
	data := new(query.AllPatientsInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func AllPatientsOutput() *query.AllPatientsOutput {
	data := new(query.AllPatientsOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
