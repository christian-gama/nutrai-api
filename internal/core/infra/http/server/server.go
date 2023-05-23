package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	http1 "github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Start starts the HTTP server.
func Start(ctx context.Context) {
	l := log.MakeWithCaller()
	l.Loading(
		"Started in %s %s",
		log.LoadingDetailColor(string(env.App.Env)),
		log.LoadingColor("environment"),
	)
	l.Loading(
		"HTTP server is running on %s",
		log.LoadingDetailColor(fmt.Sprintf("%s:%d", env.App.Host, env.App.Port)),
	)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", env.App.Port),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           http1.Router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()

	<-ctx.Done()
	if err := server.Shutdown(ctx); err != nil {
		l.Fatal(err)
	}
}
