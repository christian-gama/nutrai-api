package fake

import (
	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Message() *gpt.Message {
	data := new(gpt.Message)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	data.Role = gpt.User

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
