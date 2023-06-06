package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// AuthStrategy is the interface that defines the authentication strategy. It is used to define the
// middleware that will be used to authenticate the request.
type AuthStrategy interface {
	Middleware() middleware.Middleware
}

var (
	// AuthJwtStrategy is the authentication strategy that uses JWT.
	AuthJwtStrategy = new(authJwtStrategy)

	// AuthPublicStrategy is the authentication strategy that does not require authentication.
	AuthPublicStrategy = new(authPublicStrategy)

	// AuthApiKeyStrategy is the authentication strategy that uses API keys in the header.
	AuthApiKeyStrategy = new(authApiKeyStrategy)
)

type authJwtStrategy struct {
	middleware middleware.Middleware
}

func (s *authJwtStrategy) SetMiddleware(middleware middleware.Middleware) {
	s.middleware = middleware
}

func (s *authJwtStrategy) Middleware() middleware.Middleware {
	if s.middleware == nil {
		log.Fatal(errors.InternalServerError("authJwtStrategy middleware is not set"))
	}

	return s.middleware
}

type authPublicStrategy struct{}

func (s *authPublicStrategy) SetMiddleware(middleware middleware.Middleware) {
	log.Warn("authPublicStrategy middleware should not be set - ignoring")
}

func (s *authPublicStrategy) Middleware() middleware.Middleware {
	return nil
}

type authApiKeyStrategy struct {
	middleware middleware.Middleware
}

func (s *authApiKeyStrategy) SetMiddleware(middleware middleware.Middleware) {
	if s.middleware != nil {
		log.Warn("authApiKeyStrategy middleware is already set - overwriting")
	}

	s.middleware = middleware
}

func (s *authApiKeyStrategy) Middleware() middleware.Middleware {
	if s.middleware == nil {
		log.Fatal(errors.InternalServerError("authApiKeyStrategy middleware is not set"))
	}

	return s.middleware
}
