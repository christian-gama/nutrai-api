package middleware

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

type RecoveryStrategy interface {
	Middleware() middleware.Middleware
}

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
		panic(errors.New("recoveryAndPersistStrategy middleware is not set"))
	}

	return s.middleware
}
