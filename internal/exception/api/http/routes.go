package http

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/internal/exception/api/http/middleware"
)

func Routes() {
	router.Routes = append(router.Routes, &router.Routing{
		Middlewares: []http.Middleware{
			middleware.MakeSaveExceptionHandler(),
		},
	})
}
