package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Start starts the HTTP server.
func Start(ctx context.Context) {
	log.Loading(
		"Started in %s %s",
		log.LoadingDetailColor(string(env.App.Env)),
		log.LoadingColor("environment"),
	)
	log.Loading(
		"HTTP server is running on %s",
		log.LoadingDetailColor(fmt.Sprintf("localhost:%d", env.App.Port)),
	)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", env.App.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router.Router,
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
