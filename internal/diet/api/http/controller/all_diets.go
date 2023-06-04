package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// AllPlans is a controller to fetch all patients.
type AllPlans = controller.Controller

// NewAllPlans returns a new controller to fetch all patients.
func NewAllPlans(allPlansHandler query.AllPlansHandler) AllPlans {
	errutil.MustBeNotEmpty("query.AllPlansHandler", allPlansHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *query.AllPlansInput) {
			result, err := allPlansHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, result)
		},

		controller.Options{
			Path:     controller.JoinPath(""),
			Method:   http.MethodGet,
			IsPublic: true,
		},
	)
}
