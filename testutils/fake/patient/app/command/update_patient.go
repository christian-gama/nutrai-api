package fake

import (
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func UpdatePatientInput() *command.UpdatePatientInput {
	data := new(command.UpdatePatientInput)
	data.User = fake.User()
	data.ID = data.User.ID

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
