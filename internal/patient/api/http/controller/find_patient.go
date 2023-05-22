package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/gin-gonic/gin"
)

// FindPatient is a controller to fetch all patients.
type FindPatient = http.Controller

// NewFindPatient returns a new controller to fetch all patients.
func NewFindPatient(findPatientHandler query.FindPatientHandler) FindPatient {
	if findPatientHandler == nil {
		panic(errors.New("query.FindPatientHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *query.FindPatientInput, meta *http.Meta) {
			result, err := findPatientHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, result)
		},

		http.ControllerOptions{
			Path:     http.JoinPath(""),
			Method:   http.MethodGet,
			Params:   http.AddParams("id"),
			IsPublic: true,
		},
	)
}
