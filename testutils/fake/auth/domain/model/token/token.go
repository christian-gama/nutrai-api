package fake

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Token() *token.Token {
	data := new(token.Token)

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
