package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/gin-gonic/gin"
)

// DeleteMe is a controller to delete the current user.
type DeleteMe = http.Controller

// NewDeleteMe returns a new controller to delete the current user.
func NewDeleteMe(deleteMeHandler command.DeleteMeHandler) DeleteMe {
	if deleteMeHandler == nil {
		panic(errors.New("command.DeleteMeHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.DeleteMeInput) {
			err := deleteMeHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.NoContent(ctx)
		},

		http.ControllerOptions{
			Path:   http.JoinPath(""),
			Method: http.MethodDelete,
		},
	)
}
