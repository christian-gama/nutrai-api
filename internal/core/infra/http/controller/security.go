package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

// Security is the security type of the controller.
type Security struct {
	name       string
	middleware middleware.Middleware
}

func (s Security) Name() string {
	if s.name == "" {
		return SecurityJwt.name
	}

	return s.name
}

// String returns the string representation of the security.
func (s Security) String() string {
	return s.Name()
}

// Middleware returns the middleware of the security.
func (s Security) Middleware() middleware.Middleware {
	if s.name != "public" && s.middleware == nil {
		log.Fatalf("You must provide a middleware for security %s", s.name)
	}

	return s.middleware
}

func (s *Security) SetMiddleware(m middleware.Middleware) {
	s.middleware = m
}

var (
	// SecurityJwt will use the jwt auth middleware. It's the default security.
	SecurityJwt = &Security{
		name:       "jwt",
		middleware: nil,
	}

	// SecurityApiKey will use the api key auth middleware, which uses the APP_API_KEY secret.
	SecurityApiKey = &Security{
		name:       "api_key",
		middleware: nil,
	}

	// SecurityPublic will not use any auth middleware - it's accessible to everyone.
	SecurityPublic = &Security{
		name:       "public",
		middleware: nil,
	}
)
