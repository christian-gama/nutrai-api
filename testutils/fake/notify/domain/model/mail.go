package fake

import (
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	. "github.com/christian-gama/nutrai-api/testutils/fake"
	"github.com/go-faker/faker/v4"
)

func Mail() *mail.Mail {
	data := new(mail.Mail)
	data.To = []*value.To{
		{Name: faker.Name(), Email: faker.Email()},
	}

	err := faker.FakeData(data)
	if err != nil {
		ErrGenerating(err)
	}

	if _, err := data.Validate(); err != nil {
		ErrGenerating(err)
	}

	return data
}
