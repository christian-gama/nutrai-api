package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/metrics/http"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics = middleware.Middleware

func NewMetrics() Metrics {
	return middleware.NewMiddleware(func(ctx *gin.Context) {
		fullpath := ctx.FullPath()
		timer := prometheus.NewTimer(http.RequestsDuration.WithLabelValues(fullpath))
		http.RequestsTotal.WithLabelValues(fullpath).Inc()
		defer timer.ObserveDuration()

		defer func() {
			if r := recover(); r != nil {
				http.RequestsErrors.WithLabelValues(fullpath).Inc()
				panic(r)
			}
		}()

		ctx.Next()
	})
}
