package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// SavePlan is a controller to save a plan.
type SavePlan = controller.Controller

// NewSavePlan returns a new controller to save a plan.
func NewSavePlan(c command.SavePlanHandler) SavePlan {
	errutil.MustBeNotEmpty("command.SavePlanHandler", c)

	return controller.NewController(
		func(ctx *gin.Context, input *command.SavePlanInput) {
			err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Created(ctx, nil)
		},

		controller.Options{
			Path:   controller.JoinPath(""),
			Method: http.MethodPost,
		},
	)
}
