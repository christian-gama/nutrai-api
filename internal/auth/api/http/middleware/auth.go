package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// Auth is the middleware that handles the authentication. It will read the JWT token from
// the request header, validate it and set the user in the context. If the token is invalid, it will
// return an error.
type Auth = middleware.Middleware

// NewAuth creates a new AuthHandler.
func NewAuth(authHandler query.AuthHandler) Auth {
	errutil.MustBeNotEmpty("query.AuthHandler", authHandler)

	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			token, err := jwt.GetTokenFromHeader(ctx)
			if err != nil {
				panic(err)
			}

			authOutput, err := authHandler.Handle(ctx, &query.AuthInput{Access: token})
			if err != nil {
				panic(err)
			}

			ctxstore.SetUser(ctx,
				user.NewUser().
					SetID(authOutput.ID).
					SetEmail(authOutput.Email).
					SetName(authOutput.Name).
					SetPassword(authOutput.Password),
			)

			ctx.Next()
		},
	)
}
