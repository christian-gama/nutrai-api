package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/gin-gonic/gin"
)

// RefreshToken is a controller that handles the refreshToken of a user.
type RefreshToken = controller.Controller

// NewRefreshToken returns a new controller to handle the refreshToken of a user.
func NewRefreshToken(refreshTokenHandler service.RefreshTokenHandler) RefreshToken {
	if refreshTokenHandler == nil {
		panic(errors.New("service.RefreshTokenHandler cannot be nil"))
	}

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
			IsPublic: true,
		},
	)
}
