package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/gin-gonic/gin"
)

// UpdatePatient is a controller to update a patient.
type UpdatePatient = controller.Controller

// NewUpdatePatient returns a new controller to update a patient.
func NewUpdatePatient(updatePatientHandler command.UpdatePatientHandler) UpdatePatient {
	if updatePatientHandler == nil {
		panic(errors.New("command.UpdatePatientHandler cannot be nil"))
	}

	return controller.NewController(
		func(ctx *gin.Context, input *command.UpdatePatientInput) {
			err := updatePatientHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, nil)
		},

		controller.Options{
			Path:   controller.JoinPath(""),
			Method: http.MethodPut,
			Params: controller.AddParams("id"),
		},
	)
}
