package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// Register is a controller that handles the register of a user.
type Register = controller.Controller

// NewRegister returns a new controller to handle the register of a user.
func NewRegister(registerHandler service.RegisterHandler) Register {
	errutil.MustBeNotEmpty("service.RegisterHandler", registerHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *service.RegisterInput) {
			output, err := registerHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}

			response.Created(ctx, output)
		},

		controller.Options{
			Path:         controller.JoinPath("register"),
			Method:       http.MethodPost,
			AuthStrategy: controller.AuthPublicStrategy,
			RPM:          20,
		},
	)
}
