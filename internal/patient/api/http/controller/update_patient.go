package controller

import (
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	"github.com/gin-gonic/gin"
)

// UpdatePatient is a controller to update a patient.
type UpdatePatient = http.Controller

// NewUpdatePatient returns a new controller to update a patient.
func NewUpdatePatient(updatePatientHandler command.UpdatePatientHandler) UpdatePatient {
	if updatePatientHandler == nil {
		panic(errors.New("command.UpdatePatientHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.UpdatePatientInput, meta *http.Meta) {
			fmt.Printf("currentUser is %v", meta.CurrentUser())

			err := updatePatientHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, nil)
		},

		http.ControllerOptions{
			Path:   http.JoinPath(""),
			Method: http.MethodPut,
			Params: http.AddParams("id"),
		},
	)
}
