package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/gin-gonic/gin"
)

// UpdatePatient is a controller to update a patient.
type UpdatePatient = http.Controller

// NewUpdatePatient returns a new controller to update a patient.
func NewUpdatePatient(c command.UpdatePatientHandler) UpdatePatient {
	if c == nil {
		panic(errors.New("command.UpdatePatientHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.UpdatePatientInput) {
			err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, nil)
		},

		http.ControllerOptions{
			Path:   http.JoinPath(""),
			Method: http.MethodPut,
			Params: http.AddParams("id"),
		},
	)
}
