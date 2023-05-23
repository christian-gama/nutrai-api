package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/gin-gonic/gin"
)

// AllPatients is a controller to fetch all patients.
type AllPatients = controller.Controller

// NewAllPatients returns a new controller to fetch all patients.
func NewAllPatients(allPatientsHandler query.AllPatientsHandler) AllPatients {
	if allPatientsHandler == nil {
		panic(errors.New("query.AllPatientsHandler cannot be nil"))
	}

	return controller.NewController(
		func(ctx *gin.Context, input *query.AllPatientsInput) {
			result, err := allPatientsHandler.Handle(ctx.Request.Context(), input)
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
