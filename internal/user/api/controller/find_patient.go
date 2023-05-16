package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/user/app/query"
	"github.com/gin-gonic/gin"
)

// FindPatient is a controller to fetch all patients.
type FindPatient = http.Controller

// NewFindPatient returns a new controller to fetch all patients.
func NewFindPatient(q query.FindPatientHandler) FindPatient {
	if q == nil {
		panic(errors.New("query.FindPatientHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *query.FindPatientInput) {
			result, err := q.Handle(ctx.Request.Context(), input)
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
