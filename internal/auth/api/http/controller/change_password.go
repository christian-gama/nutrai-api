package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// ChangePassword is a controller to change the current user's password.
type ChangePassword = controller.Controller

// NewChangePassword returns a new controller to change the current user's password.
func NewChangePassword(changePasswordHandler command.ChangePasswordHandler) ChangePassword {
	errutil.MustBeNotEmpty("command.ChangePasswordHandler", changePasswordHandler)

	m := &changePasswordHandlerImpl{
		changePasswordHandler: changePasswordHandler,
	}

	return controller.NewController(
		m.Handle,
		controller.Options{
			Path:   controller.JoinPath(""),
			Method: http.MethodPatch,
		},
	)
}

type changePasswordHandlerImpl struct {
	changePasswordHandler command.ChangePasswordHandler
}

func (h *changePasswordHandlerImpl) Handle(
	ctx *gin.Context,
	input *command.ChangePasswordInput,
) {
	err := h.changePasswordHandler.Handle(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	response.Ok(ctx, nil)
}
