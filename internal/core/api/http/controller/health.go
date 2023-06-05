package controller

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/gin-gonic/gin"
)

// Health is the controller that handles the health check of the API.
type Health = controller.Controller

// NewHealth creates a new health controller.
func NewHealth() Health {
	return controller.NewController(
		func(ctx *gin.Context, _ *any) {
			response.Ok(ctx, time.Now())
		},

		controller.Options{
			Security: controller.SecurityApiKey,
			Path:     controller.JoinPath("health"),
			Method:   http.MethodGet,
			RPM:      300,
		},
	)
}
