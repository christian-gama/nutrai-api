package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/gin-gonic/gin"
)

// SavePatient is a controller to save a patient.
type SavePatient = controller.Controller

// NewSavePatient returns a new controller to save a patient.
func NewSavePatient(c command.SavePatientHandler) SavePatient {
	if c == nil {
		panic(errors.New("command.SavePatientHandler cannot be nil"))
	}

	return controller.NewController(
		func(ctx *gin.Context, input *command.SavePatientInput) {
			err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Created(ctx, nil)
		},

		controller.Options{
			Path:   controller.JoinPath(""),
			Method: http.MethodPost,
		},
	)
}
