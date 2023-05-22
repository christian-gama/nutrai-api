package router

import "github.com/christian-gama/nutrai-api/internal/core/domain/logger"

// Register registers all the routes for a group.
func Register(log logger.Logger) {
	log.Info("Registering routes")

	group := Router.Group("api")

	for _, route := range Routes {
		route.Register(group)
	}
}
