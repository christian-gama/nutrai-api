package internal

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth"
	"github.com/christian-gama/nutrai-api/internal/core"
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/internal/diet"
	"github.com/christian-gama/nutrai-api/internal/exception"
	"github.com/christian-gama/nutrai-api/internal/notify"
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.NewLoader(envFile).Load()
	log.SugaredLogger = log.New()
	router.SetupRouter()

	// Order matters.
	module.Init(exception.Init)
	module.Init(auth.Init)
	module.Init(core.Init)
	module.Init(patient.Init)
	module.Init(diet.Init)
	module.Init(notify.Init)
}
