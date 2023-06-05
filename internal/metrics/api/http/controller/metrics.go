package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metric = controller.Controller

func NewMetric() Metric {
	return controller.NewController(
		func(ctx *gin.Context, _ *any) {
			promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
		},

		controller.Options{
			IsPublic: true,
			Path:     controller.JoinPath("metrics"),
			Method:   http.MethodGet,
			RPM:      300,
		},
	)
}
