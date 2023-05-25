package suite

import (
	"github.com/christian-gama/nutrai-api/config/env"
	authValidation "github.com/christian-gama/nutrai-api/internal/auth/infra/validation"
	patientValidation "github.com/christian-gama/nutrai-api/internal/patient/infra/validation"
	"github.com/christian-gama/nutrai-api/testutils/faker"
)

func init() {
	faker.Init()

	authValidation.Register()
	patientValidation.Register()

	env.NewLoader(".env.test").Load()
}
