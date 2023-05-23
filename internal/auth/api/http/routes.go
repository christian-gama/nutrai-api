package http

import (
	. "github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
)

// Routes registers the routes for user module.
func Routes() {
	router.Routes = append(router.Routes, &router.Routing{
		Group: "/v1/auth",
		Routes: []*router.Route{
			{Controller: MakeLogin()},
			{Controller: MakeRegister()},
			{Controller: MakeRefreshToken()},
		},
	})

	router.Routes = append(router.Routes, &router.Routing{
		Group: "/v1/users",
		Routes: []*router.Route{
			{Controller: MakeDeleteUser()},
			{Controller: MakeChangePassword()},
		},
	})
}
