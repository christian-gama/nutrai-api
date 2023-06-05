package middleware

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/gin-gonic/gin"
)

type ApiKeyAuth = middleware.Middleware

func NewApiKeyAuth() ApiKeyAuth {
	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			if key, err := http.GetAuthorizationHeader(ctx.Request); err != nil {
				panic(err)
			} else if key != env.App.ApiKey {
				panic("invalid api key")
			}

			ctx.Next()
		},
	)
}
