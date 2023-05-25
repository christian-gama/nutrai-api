package middleware

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
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
	if authHandler == nil {
		panic(errors.New("query.AuthHandler is required"))
	}

	return middleware.NewMiddleware(
		func(ctx *gin.Context) {
			token, err := jwt.GetTokenFromHeader(ctx)
			if err != nil {
				handleUnauthorizedError(err)
			}

			u, err := authHandler.Handle(ctx, &query.AuthInput{Access: token})
			if err != nil {
				handleUnauthorizedError(err)
			}

			store.SetUser(ctx,
				user.NewUser().
					SetID(u.ID).
					SetEmail(u.Email).
					SetName(u.Name).
					SetPassword(u.Password),
			)

			ctx.Next()
		},
	)
}

func handleUnauthorizedError(err error) {
	var unauthorizedErr *errutil.ErrUnauthorized
	if errors.As(err, &unauthorizedErr) {
		panic(unauthorizedErr)
	}

	panic(errutil.Unauthorized(err.Error()))
}
