package fake

import (
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	"github.com/christian-gama/nutrai-api/testutils/fake"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
	"github.com/go-faker/faker/v4"
)

func Patient() *patient.Patient {
	data := new(patient.Patient)
	data.User = userFake.User()
	data.ID = data.User.ID

	err := faker.FakeData(data)
	if err != nil {
		fake.ErrGenerating(err)
	}

	if err := data.Validate(); err != nil {
		fake.ErrGenerating(err)
	}

	return data
}
