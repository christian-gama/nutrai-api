package router

import (
	"fmt"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/unit"
	_cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// logging returns a gin middleware that logs the request.
func logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		duration := bench.Duration(ctx.Next)

		logLevel(ctx.Writer.Status(), duration)(
			"%-6s | %-5s | %6s | %s",
			ctx.Request.Method,
			statusColor(ctx.Writer.Status()),
			duration.Truncate(time.Millisecond),
			ctx.Request.URL.Path,
		)
	}
}

// cors returns a gin middleware that enables CORS.
func cors() gin.HandlerFunc {
	return _cors.New(_cors.Config{
		AllowFiles:    true,
		AllowOrigins:  env.App.AllowedOrigins,
		AllowWildcard: true,
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods: []string{
			http.MethodGet.String(),
			http.MethodPost.String(),
			http.MethodPut.String(),
			http.MethodDelete.String(),
			http.MethodPatch.String(),
		},
	})
}

// limitBodySize returns a gin middleware that limits the request body size.
func limitBodySize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const maxBodySize = 3 * unit.Megabyte

		if ctx.Request.ContentLength > maxBodySize {
			response.BadRequest(
				ctx,
				errors.Invalid(
					"body",
					fmt.Sprintf("body size must be less than %d bytes", maxBodySize),
				),
			)
			return
		}

		ctx.Next()
	}
}

// RateLimiter returns a gin middleware that limits the request rate.
func RateLimiter(limit int, duration time.Duration) gin.HandlerFunc {
	if limit == 0 || !env.Config.EnableRateLimit {
		return func(ctx *gin.Context) {
			ctx.Next()
		}
	}

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  duration,
		Limit: uint(limit),
	})

	handler := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, info ratelimit.Info) {
			response.TooManyRequests(
				ctx,
				errors.TooManyRequests(),
			)
		},

		KeyFunc: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},
	})

	return handler
}
