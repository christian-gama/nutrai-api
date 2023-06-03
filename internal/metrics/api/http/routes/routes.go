package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/internal/metrics/api/http/controller"
)

// Register registers the routes for  module.
func Register() {
	routes.Root().
		SetController(controller.MakeMetric())
}
