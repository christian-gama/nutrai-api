package router

import (
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/gin-gonic/gin"
)

// Router is the global router of the application.
var Router *gin.Engine

// SetupRouter sets the mode of the router and returns a new router.
// It will also set up the global middlewares.
func SetupRouter() {
	if env.IsDevelopment && env.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(cors())
	r.Use(logging())
	r.Use(limitBodySize())
	r.Use(RateLimiter(env.Config.GlobalRateLimit, time.Minute))

	Router = r
}
