package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// Logout is a controller to delete the current user.
type Logout = controller.Controller

// NewLogout returns a new controller to delete the current user.
func NewLogout(logoutHandler command.LogoutHandler) Logout {
	errutil.MustBeNotEmpty("command.LogoutHandler", logoutHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *command.LogoutInput) {
			err := logoutHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, nil)
		},

		controller.Options{
			Path:   controller.JoinPath("logout"),
			Method: http.MethodPost,
		},
	)
}
