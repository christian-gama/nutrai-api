package routes

import (
	"github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
)

// Register registers the routes for  module.
func Register() {
	routes.Api("/v1/auth").
		SetController(controller.MakeLogin()).
		SetController(controller.MakeRegister()).
		SetController(controller.MakeRefreshToken()).
		SetController(controller.MakeLogout())

	routes.Api("/v1/users").
		SetController(controller.MakeDeleteUser()).
		SetController(controller.MakeChangePassword())
}
