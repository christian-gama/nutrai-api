package internal

import (
	"context"

	// Initialize custom validation aliases.
	"github.com/christian-gama/nutrai-api/internal/shared/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/env"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/server"
	_ "github.com/christian-gama/nutrai-api/internal/shared/infra/validation"
	"github.com/christian-gama/nutrai-api/internal/user"
)

// Bootstrap is the main function that starts the application.
func Bootstrap(ctx context.Context, log logger.Logger, envFile string) {
	env.Load(envFile)
	log.Infof("Booting the application")

	r := router.New()
	routing.Register(r.Group("/api"),
		user.Routes(),
	)

	server.Start(ctx, r, log)
}
