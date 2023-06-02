package middleware

import (
	"fmt"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/gin-gonic/gin"
)

type RateLimiter = middleware.Middleware

func NewRateLimiter(store ratelimit.Store) RateLimiter {
	handler := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: func(ctx *gin.Context, info ratelimit.Info) {
			response.TooManyRequests(ctx, errors.TooManyRequests())
		},

		KeyFunc: func(ctx *gin.Context) string {
			return ctx.ClientIP()
		},

		BeforeResponse: func(ctx *gin.Context, info ratelimit.Info) {
			ctx.Header("X-RateLimit-Limit", fmt.Sprint(info.Limit))
			ctx.Header("X-RateLimit-Remaining", fmt.Sprint(info.RemainingHits))
			ctx.Header("X-RateLimit-Reset", fmt.Sprint(info.ResetTime.UTC().Format(time.RFC3339)))
		},
	})

	return middleware.NewMiddleware(func(ctx *gin.Context) {
		handler(ctx)
		ctx.Next()
	})
}
