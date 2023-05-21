package internal

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router"
)

// Bootstrap is responsible for booting up the application.
func Bootstrap(envFile string) {
	env.Load(envFile)
	router.Engine = router.SetupEngine()
	Routing()
	Workers()
}
