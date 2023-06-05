package middleware

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/gin-gonic/gin"
)

type ApiKey = middleware.Middleware

func NewApiKey() ApiKey {
	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			if _, err := http.CheckAuthorizationHeader(ctx.Request, env.App.ApiKey); err != nil {
				panic(err)
			}

			ctx.Next()
		},
	)
}
