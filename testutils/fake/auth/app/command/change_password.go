package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func ChangePasswordInput() *command.ChangePasswordInput {
	data := new(command.ChangePasswordInput)
	data.User = fake.User()

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
