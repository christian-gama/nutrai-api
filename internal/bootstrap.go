package internal

import (
	"context"

	// Initialize custom validation aliases.
	"github.com/christian-gama/nutrai-api/internal/shared/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/server"
	_ "github.com/christian-gama/nutrai-api/internal/shared/infra/validation"
)

// Bootstrap is the main function that starts the application.
func Bootstrap(ctx context.Context, log logger.Logger, envFile string) {
	env.Load(envFile)
	log.Infof("Booting the application")

	r := Routes()

	server.Start(ctx, r, log)
}
