package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// LogoutAll is a controller to delete the current user.
type LogoutAll = controller.Controller

// NewLogoutAll returns a new controller to delete the current user.
func NewLogoutAll(logoutAllHandler command.LogoutAllHandler) LogoutAll {
	errutil.MustBeNotEmpty("command.LogoutAllHandler", logoutAllHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *command.LogoutAllInput) {
			err := logoutAllHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, nil)
		},

		controller.Options{
			Path:   controller.JoinPath("logout-all"),
			Method: http.MethodPost,
		},
	)
}
