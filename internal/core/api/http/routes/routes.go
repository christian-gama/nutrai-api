package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
)

// Register registers the routes for  module.
func Register() {
	routes.Internal().
		SetController(controller.MakeHealth())
}
