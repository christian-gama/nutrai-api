package exception

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/exception/api/middleware"
)

func Routes() *routing.Routing {
	return &routing.Routing{
		Middlewares: []http.Middleware{
			middleware.MakeSaveExceptionHandler(),
		},
	}
}
