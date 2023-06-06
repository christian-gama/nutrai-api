package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/internal/metrics/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/metrics/api/http/middleware"
)

// Register registers the routes for  module.
func Register() {
	routes.SetGlobalMiddleware(middleware.MakeMetrics())

	routes.Internal().
		SetController(controller.MakeMetric())
}
