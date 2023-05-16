package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router"
)

// Start starts the HTTP server.
func Start(ctx context.Context) {
	log.WithCaller.Infof("Started in %s environment", env.App.Env)
	log.WithCaller.Infof("Server is running on %s:%d", env.App.Host, env.App.Port)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", env.App.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router.Engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.WithCaller.Fatal(err)
		}
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		log.WithCaller.Fatal(err)
	}
}
