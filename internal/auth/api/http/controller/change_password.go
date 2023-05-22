package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/gin-gonic/gin"
)

// ChangePassword is a controller to change the current user's password.
type ChangePassword = http.Controller

// NewChangePassword returns a new controller to change the current user's password.
func NewChangePassword(changePasswordHandler command.ChangePasswordHandler) ChangePassword {
	if changePasswordHandler == nil {
		panic(errors.New("command.ChangePasswordHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.ChangePasswordInput) {
			err := changePasswordHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, nil)
		},

		http.ControllerOptions{
			Path:   http.JoinPath(""),
			Method: http.MethodPatch,
		},
	)
}
