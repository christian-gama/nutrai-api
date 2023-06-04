package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// DeletePlan is a controller to delete the current user.
type DeletePlan = controller.Controller

// NewDeletePlan returns a new controller to delete the current user.
func NewDeletePlan(deletePlanHandler command.DeletePlanHandler) DeletePlan {
	errutil.MustBeNotEmpty("command.DeletePlanHandler", deletePlanHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *command.DeletePlanInput) {
			err := deletePlanHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.NoContent(ctx)
		},

		controller.Options{
			Path:   controller.JoinPath(""),
			Method: http.MethodDelete,
			Params: controller.AddParams("id"),
		},
	)
}
