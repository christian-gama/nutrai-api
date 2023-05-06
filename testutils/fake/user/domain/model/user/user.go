package fake

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func User() *user.User {
	data := new(user.User)

	err := faker.FakeData(data)
	if err != nil {
		panic(fmt.Errorf("error while generating fake User: %w", err))
	}

	if err := data.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake User: %w", err))
	}

	return data
}
