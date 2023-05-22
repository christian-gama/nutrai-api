package internal

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth"
	"github.com/christian-gama/nutrai-api/internal/core"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/internal/diet"
	"github.com/christian-gama/nutrai-api/internal/exception"
	"github.com/christian-gama/nutrai-api/internal/patient"
)

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.Load(envFile)
	log := log.MakeWithCaller()

	core.Init(log)
	exception.Init(log)
	auth.Init(log)
	patient.Init(log)
	diet.Init(log)

	router.Register(log)
}
