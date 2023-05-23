package router

import "github.com/christian-gama/nutrai-api/internal/core/infra/log"

// Register registers all the routes for a group.
func Register() {
	log := log.MakeWithCaller()
	log.Loading("Registering routes")

	group := Router.Group("api")

	for _, route := range Routes {
		route.Register(group)
	}
}
