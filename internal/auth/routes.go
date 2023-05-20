package auth

import (
	"github.com/christian-gama/nutrai-api/internal/auth/api/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/router/routing"
)

// Routes registers the routes for user module.
func Routes() *routing.Routing {
	return &routing.Routing{
		Group: "/v1/auth",
		Routes: []*routing.Route{
			{Controller: controller.MakeLogin()},
			{Controller: controller.MakeRegister()},
		},
	}
}
