package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func FindPatientInput() *query.FindPatientInput {
	data := new(query.FindPatientInput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}

func FindPatientOutput() *query.FindPatientOutput {
	data := new(query.FindPatientOutput)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
