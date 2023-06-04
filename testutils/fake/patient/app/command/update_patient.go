package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func UpdatePatientInput() *command.UpdatePatientInput {
	data := new(command.UpdatePatientInput)
	data.User = fake.User()
	data.Allergies = []value.Allergy{
		value.Allergy(faker.Name()),
		value.Allergy(faker.Name()),
	}

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
