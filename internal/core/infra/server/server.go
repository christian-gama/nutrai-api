package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router"
)

// Start starts the HTTP server.
func Start(ctx context.Context, log logger.Logger) {
	log.Infof("Started in %s environment", env.App.Env)
	log.Infof("Server is running on %s:%d", env.App.Host, env.App.Port)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", env.App.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
