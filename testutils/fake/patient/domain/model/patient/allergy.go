package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Allergy() *patient.Allergy {
	data := new(patient.Allergy)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
