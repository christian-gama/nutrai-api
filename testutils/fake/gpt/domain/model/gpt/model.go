package fake

import (
	"github.com/christian-gama/nutrai-api/internal/gpt/domain/model/gpt"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Model() *gpt.Model {
	data := new(gpt.Model)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
