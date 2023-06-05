package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// Login is a controller that handles the login of a user.
type Login = controller.Controller

// NewLogin returns a new controller to handle the login of a user.
func NewLogin(loginHandler service.LoginHandler) Login {
	errutil.MustBeNotEmpty("service.LoginHandler", loginHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *service.LoginInput) {
			output, err := loginHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, output)
		},

		controller.Options{
			Path:     controller.JoinPath("login"),
			Method:   http.MethodPost,
			Security: controller.SecurityPublic,
		},
	)
}
