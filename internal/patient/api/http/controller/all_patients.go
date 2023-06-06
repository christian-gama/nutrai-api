package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// AllPatients is a controller to fetch all patients.
type AllPatients = controller.Controller

// NewAllPatients returns a new controller to fetch all patients.
func NewAllPatients(allPatientsHandler query.AllPatientsHandler) AllPatients {
	errutil.MustBeNotEmpty("query.AllPatientsHandler", allPatientsHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *query.AllPatientsInput) {
			result, err := allPatientsHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, result)
		},

		controller.Options{
			Path:         controller.JoinPath(""),
			Method:       http.MethodGet,
			AuthStrategy: controller.AuthPublicStrategy,
		},
	)
}
