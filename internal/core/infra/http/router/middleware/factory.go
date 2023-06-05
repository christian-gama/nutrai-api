// Internal middlewares for the routes. It's important to note that the middlewares here are
// different from the middlewares from the api/http/middlewares, because the middlewares here
// are used by the routes internally.

package middleware

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
)

func MakeLogging() Logging {
	return NewLogging()
}

func MakeCors() Cors {
	return NewCors()
}

func MakeLimitBodySize() LimitBodySize {
	return NewLimitBodySize()
}

func MakeRateLimiter(rpm int) RateLimiter {
	return NewRateLimiter(ratelimit.RedisStore(&ratelimit.RedisOptions{
		Rate:        time.Minute,
		Limit:       uint(rpm),
		RedisClient: conn.GetRedis(),
	}))
}

func MakeRecovery() Recovery {
	return NewRecovery()
}
