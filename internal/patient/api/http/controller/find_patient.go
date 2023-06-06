package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// FindPatient is a controller to fetch all patients.
type FindPatient = controller.Controller

// NewFindPatient returns a new controller to fetch all patients.
func NewFindPatient(findPatientHandler query.FindPatientHandler) FindPatient {
	errutil.MustBeNotEmpty("query.FindPatientHandler", findPatientHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *query.FindPatientInput) {
			result, err := findPatientHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, result)
		},

		controller.Options{
			Path:         controller.JoinPath(""),
			Method:       http.MethodGet,
			Params:       controller.AddParams("id"),
			AuthStrategy: controller.AuthPublicStrategy,
		},
	)
}
