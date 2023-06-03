package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

var recoveryMiddleware middleware.Middleware

type Recovery = middleware.Middleware

// NewRecovery creates a new RecoveryHandler.
func NewRecovery() Recovery {
	return recoveryMiddleware
}

func SetRecoveryMiddleware(middleware middleware.Middleware) {
	if recoveryMiddleware != nil {
		log.Warn("recovery middleware was already set - overriding")
	}

	recoveryMiddleware = middleware
}
