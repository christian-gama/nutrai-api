package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

type AuthApiKey = middleware.Middleware

func NewAuthApiKey(AuthApiKeyHandler query.AuthApiKeyHandler) AuthApiKey {
	errutil.MustBeNotEmpty("query.AuthApiKeyHandler", AuthApiKeyHandler)

	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			key, err := http.GetAuthorizationHeader(ctx.Request)
			if err != nil {
				panic(err)
			}

			_, err = AuthApiKeyHandler.Handle(ctx, &query.AuthApiKeyInput{Key: key})
			if err != nil {
				panic(err)
			}

			ctx.Next()
		},
	)
}
