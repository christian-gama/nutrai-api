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
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.NewLoader(envFile).Load()
	log := log.MakeWithCaller()

	// Order matters.
	module.Initialize(log, exception.Init)
	module.Initialize(log, auth.Init)
	module.Initialize(log, core.Init)
	module.Initialize(log, patient.Init)
	module.Initialize(log, diet.Init)

	router.Register(log)
}
