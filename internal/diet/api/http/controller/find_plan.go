package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// FindPlan is a controller to fetch all patients.
type FindPlan = controller.Controller

// NewFindPlan returns a new controller to fetch all patients.
func NewFindPlan(findPlanHandler query.FindPlanHandler) FindPlan {
	errutil.MustBeNotEmpty("query.FindPlanHandler", findPlanHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *query.FindPlanInput) {
			result, err := findPlanHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, result)
		},

		controller.Options{
			Path:         controller.JoinPath(""),
			Method:       http.MethodGet,
			Params:       controller.AddParams("id"),
			AuthStrategy: controller.AuthJwtStrategy,
		},
	)
}
