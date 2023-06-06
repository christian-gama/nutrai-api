package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// RecoveryStrategy is the strategy responsible for setting the middleware for recoverying
// from panics.
type RecoveryStrategy interface {
	Middleware() middleware.Middleware
}

// RecoveryAndPersistStrategy is the strategy responsible for setting the middleware for recoverying
// from panics and persisting the errors in a database.
var RecoveryAndPersistStrategy = new(recoveryAndPersistStrategy)

type recoveryAndPersistStrategy struct {
	middleware middleware.Middleware
}

func (s *recoveryAndPersistStrategy) SetMiddleware(middleware middleware.Middleware) {
	if s.middleware != nil {
		log.Warn("recoveryAndPersistStrategy middleware is already set - overwriting")
	}

	s.middleware = middleware
}

func (s *recoveryAndPersistStrategy) Middleware() middleware.Middleware {
	if s.middleware == nil {
		log.Fatal(errors.InternalServerError("recoveryAndPersistStrategy middleware is not set"))
	}

	return s.middleware
}
