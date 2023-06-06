package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// JwtAuth is the middleware that handles the authentication. It will read the JWT token from
// the request header, validate it and set the user in the context. If the token is invalid, it will
// return an error.
type JwtAuth = middleware.Middleware

// NewJwtAuth creates a new AuthHandler.
func NewJwtAuth(jwtAuthHandler query.JwtAuthHandler) JwtAuth {
	errutil.MustBeNotEmpty("query.AuthHandler", jwtAuthHandler)

	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			token, err := http.GetAuthorizationHeader(ctx.Request)
			if err != nil {
				panic(err)
			}

			authOutput, err := jwtAuthHandler.Handle(
				ctx,
				&query.JwtAuthInput{Access: value.Token(token)},
			)
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
