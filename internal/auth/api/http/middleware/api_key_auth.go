package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

type ApiKeyAuth = middleware.Middleware

func NewApiKeyAuth(apiKeyAuthHandler query.ApiKeyAuthHandler) ApiKeyAuth {
	errutil.MustBeNotEmpty("query.ApiKeyAuthHandler", apiKeyAuthHandler)

	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			key, err := http.GetAuthorizationHeader(ctx.Request)
			if err != nil {
				panic(err)
			}

			_, err = apiKeyAuthHandler.Handle(ctx, &query.ApiKeyAuthInput{Key: key})
			if err != nil {
				panic(err)
			}

			ctx.Next()
		},
	)
}
