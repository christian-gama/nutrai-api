package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/http"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/gin-gonic/gin"
)

// SavePatient is a controller to save a patient.
type SavePatient = http.Controller

// NewSavePatient returns a new controller to save a patient.
func NewSavePatient(c command.SavePatientHandler) SavePatient {
	if c == nil {
		panic(errors.New("command.SavePatientHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.SavePatientInput) {
			err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Created(ctx, nil)
		},

		http.ControllerOptions{
			Path:     "/",
			Method:   http.MethodPost,
			IsPublic: true,
		},
	)
}
