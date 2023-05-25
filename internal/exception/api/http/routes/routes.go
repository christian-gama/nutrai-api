package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/internal/exception/api/http/middleware"
)

// Register registers the routes for this module.
func Register() {
	routes.SetGlobalMiddleware(middleware.MakeSaveException())
}
