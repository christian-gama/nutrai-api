package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

var authMiddleware middleware.Middleware

type Auth = middleware.Middleware

// NewAuth creates a new AuthHandler.
func NewAuth() Auth {
	return authMiddleware
}

func SetAuthMiddleware(middleware middleware.Middleware) {
	if authMiddleware != nil {
		log.Warn("auth middleware was already set - overriding")
	}

	authMiddleware = middleware
}
