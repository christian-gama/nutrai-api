package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/gin-gonic/gin"
)

// Register is a controller that handles the register of a user.
type Register = controller.Controller

// NewRegister returns a new controller to handle the register of a user.
func NewRegister(registerHandler service.RegisterHandler) Register {
	if registerHandler == nil {
		panic(errors.New("service.RegisterHandler cannot be nil"))
	}

	return controller.NewController(
		func(ctx *gin.Context, input *service.RegisterInput) {
			output, err := registerHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Created(ctx, output)
		},

		controller.Options{
			Path:     controller.JoinPath("register"),
			Method:   http.MethodPost,
			IsPublic: true,
		},
	)
}
