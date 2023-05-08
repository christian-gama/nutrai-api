package fake

import (
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func User() *user.User {
	data := new(user.User)

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating("user", err)
	}

	if err := data.Validate(); err != nil {
		fake.ErrGenerating("user", err)
	}

	return data
}
