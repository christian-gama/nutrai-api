package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

var MetricsStrategy = new(metricsStrategy)

type metricsStrategy struct {
	middleware middleware.Middleware
}

func (m *metricsStrategy) SetMiddleware(middleware middleware.Middleware) {
	if m.middleware != nil {
		log.Warn("metricsStrategy middleware is already set - overwriting")
	}

	m.middleware = middleware
}

func (m *metricsStrategy) Middleware() middleware.Middleware {
	if m.middleware == nil {
		log.Warn("metricsStrategy middleware is not set")
	}

	return m.middleware
}
