package internal

import (
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router/routing"
	"github.com/christian-gama/nutrai-api/internal/user"
	"github.com/gin-gonic/gin"
)

// LoadEngine registers all the routes.
func LoadEngine() *gin.Engine {
	engine := router.New()

	RegisterApiRoutes(engine)

	return engine
}

// RegisterApiRoutes registers all the routes for API. The modules should be
// added here.
func RegisterApiRoutes(engine *gin.Engine) ([]*routing.Routing, string) {
	group, routes := routing.Register(engine.Group("/api"),
		// Add all routes for modules here (order matters):
		user.Routes(),
	)

	return routes, group.BasePath()
}
