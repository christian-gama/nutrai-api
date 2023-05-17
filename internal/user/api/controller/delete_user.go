package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/gin-gonic/gin"
)

// DeleteUser is a controller to delete a user.
type DeleteUser = http.Controller

// NewDeleteUser returns a new controller to delete a user.
func NewDeleteUser(c command.DeleteUserHandler) DeleteUser {
	if c == nil {
		panic(errors.New("command.DeleteUserHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *command.DeleteUserInput) {
			err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.NoContent(ctx)
		},

		http.ControllerOptions{
			Path:   http.JoinPath(""),
			Method: http.MethodDelete,
			Params: http.AddParams("id"),
		},
	)
}
