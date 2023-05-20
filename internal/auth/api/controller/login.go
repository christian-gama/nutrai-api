package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/gin-gonic/gin"
)

// Login is a controller that handles the login of a user.
type Login = http.Controller

// NewLogin returns a new controller to handle the login of a user.
func NewLogin(c service.LoginHandler) Login {
	if c == nil {
		panic(errors.New("command.LoginHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.LoginInput) {
			output, err := c.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, output)
		},

		http.ControllerOptions{
			Path:     http.JoinPath("login"),
			Method:   http.MethodPost,
			IsPublic: true,
		},
	)
}
