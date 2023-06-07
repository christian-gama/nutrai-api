package router

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/middleware"
	"github.com/gin-gonic/gin"
)

// Router is the global router of the application.
var (
	Router *gin.Engine
)

// SetupRouter sets the mode of the router and returns a new router.
// It will also set up the global middlewares.
func SetupRouter() {
	if env.IsDevelopment && env.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.New()
	Router.Use(middleware.RecoveryAndPersistStrategy.Middleware().Handle)
	Router.Use(middleware.MetricsStrategy.Middleware().Handle)
	Router.Use(middleware.MakeCors().Handle)
	Router.Use(middleware.MakeLogging().Handle)
	Router.Use(middleware.MakeLimitBodySize().Handle)
}
