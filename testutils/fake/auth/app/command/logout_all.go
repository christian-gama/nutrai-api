package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func LogoutAllInput() *command.LogoutAllInput {
	data := new(command.LogoutAllInput)
	data.User = fake.User()

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	return data
}
