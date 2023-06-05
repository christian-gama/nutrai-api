package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// RefreshToken is a controller that handles the refreshToken of a user.
type RefreshToken = controller.Controller

// NewRefreshToken returns a new controller to handle the refreshToken of a user.
func NewRefreshToken(refreshTokenHandler service.RefreshTokenHandler) RefreshToken {
	errutil.MustBeNotEmpty("service.RefreshTokenHandler", refreshTokenHandler)

	return controller.NewController(
		func(ctx *gin.Context, input *service.RefreshTokenInput) {
			output, err := refreshTokenHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			response.Ok(ctx, output)
		},

		controller.Options{
			Path:     controller.JoinPath("refresh"),
			Method:   http.MethodPost,
			Security: controller.SecurityPublic,
		},
	)
}
