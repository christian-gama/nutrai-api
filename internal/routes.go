package internal

import (
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/user"
	"github.com/gin-gonic/gin"
)

func ApiRoutes() (*gin.Engine, string, []*routing.Routing) {
	engine := router.New()

	group, routes := routing.Register(engine.Group("/api"),
		// Register here the modules routes.
		user.Routes(),
	)

	return engine, group.BasePath(), routes
}
