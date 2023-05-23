package internal

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth"
	"github.com/christian-gama/nutrai-api/internal/core"
	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/diet"
	"github.com/christian-gama/nutrai-api/internal/exception"
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.NewLoader(envFile).Load()

	// Order matters.
	module.Initialize(exception.Init)
	module.Initialize(auth.Init)
	module.Initialize(core.Init)
	module.Initialize(patient.Init)
	module.Initialize(diet.Init)

	router.Register()
}
