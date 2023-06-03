package suite

import (
	"github.com/christian-gama/nutrai-api/config/env"
	authValidation "github.com/christian-gama/nutrai-api/internal/auth/infra/validation"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	patientValidation "github.com/christian-gama/nutrai-api/internal/patient/infra/validation"
	"github.com/christian-gama/nutrai-api/testutils/faker"
	"go.uber.org/zap"
)

func init() {
	env.NewLoader(".env.test").Load()

	log.SugaredLogger = zap.NewNop().Sugar()
	faker.Init()

	authValidation.Register()
	patientValidation.Register()
}
