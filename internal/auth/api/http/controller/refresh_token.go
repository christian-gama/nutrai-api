package controller

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/gin-gonic/gin"
)

// RefreshToken is a controller that handles the refreshToken of a user.
type RefreshToken = http.Controller

// NewRefreshToken returns a new controller to handle the refreshToken of a user.
func NewRefreshToken(refreshTokenHandler service.RefreshTokenHandler) RefreshToken {
	if refreshTokenHandler == nil {
		panic(errors.New("service.RefreshTokenHandler cannot be nil"))
	}

	return http.NewController(
		func(ctx *gin.Context, input *service.RefreshTokenInput, meta *http.Meta) {
			output, err := refreshTokenHandler.Handle(ctx.Request.Context(), input)
			if err != nil {
				panic(err)
			}
			http.Ok(ctx, output)
		},

		http.ControllerOptions{
			Path:     http.JoinPath("refresh"),
			Method:   http.MethodPost,
			IsPublic: true,
		},
	)
}
