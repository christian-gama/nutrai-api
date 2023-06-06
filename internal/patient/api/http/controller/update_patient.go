package controller

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// UpdatePatient is a controller to update a patient.
type UpdatePatient = controller.Controller

// NewUpdatePatient returns a new controller to update a patient.
func NewUpdatePatient(updatePatientHandler command.UpdatePatientHandler) UpdatePatient {
	errutil.MustBeNotEmpty("command.UpdatePatientHandler", updatePatientHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *command.UpdatePatientInput) {
			err := updatePatientHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, nil)
		},

		controller.Options{
			Path:         controller.JoinPath(""),
			Method:       http.MethodPut,
			AuthStrategy: controller.AuthJwtStrategy,
		},
	)
}
