package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/gin-gonic/gin"
)

// AllPatients is a controller to fetch all patients.
type AllPatients = http.Controller

// NewAllPatients returns a new controller to fetch all patients.
func NewAllPatients(allPatientsHandler query.AllPatientsHandler) AllPatients {
	if allPatientsHandler == nil {
		panic(errors.New("query.AllPatientsHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *query.AllPatientsInput) {
			result, err := allPatientsHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, result)
		},

		http.ControllerOptions{
			Path:     http.JoinPath(""),
			Method:   http.MethodGet,
			IsPublic: true,
		},
	)
}
