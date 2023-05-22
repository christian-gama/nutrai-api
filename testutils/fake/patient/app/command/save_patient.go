package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func SavePatientInput() *command.SavePatientInput {
	data := new(command.SavePatientInput)
	data.User = fake.User()

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
