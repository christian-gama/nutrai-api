package router

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	"github.com/gin-gonic/gin"
)

// SetupEngine sets the mode of the router and returns a new router.
// It will also set up the global middlewares.
func SetupEngine() *gin.Engine {
	if env.App.Env == env.Development && env.Config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(Logging())
	r.Use(Cors())
	r.Use(Content())
	r.Use(LimitBodySize())
	r.Use(RateLimiter(env.Config.GlobalRateLimit, time.Minute))

	return r
}

var Engine *gin.Engine
