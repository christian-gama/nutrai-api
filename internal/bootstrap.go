package internal

import (
	"context"

	// Initialize custom validation aliases.
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/server"
	_ "github.com/christian-gama/nutrai-api/internal/core/infra/validation"
)

// Bootstrap is the main function that starts the application.
func Bootstrap(ctx context.Context, log logger.Logger, envFile string) {
	env.Load(envFile)
	log.Infof("Booting the application")

	engine := LoadEngine()

	server.Start(ctx, engine, log)
}
