package middleware

import (
	"errors"

	_jwt "github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/gin-gonic/gin"
)

// AuthHandler is the middleware that handles the authentication.
type AuthHandler = http.Middleware

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(verifier _jwt.Verifier, userRepo repo.User) AuthHandler {
	if verifier == nil {
		panic(errors.New("jwt.Verifier is required"))
	}

	if userRepo == nil {
		panic(errors.New("repo.user is required"))
	}

	return http.NewMiddleware(
		func(ctx *gin.Context) {
			token, err := jwt.GetTokenFromHeader(ctx)
			if err != nil {
				panic(err)
			}

			claims, err := verifier.Verify(token)
			if err != nil {
				panic(err)
			}

			u, err := userRepo.FindByEmail(
				ctx.Request.Context(),
				repo.FindByEmailUserInput{Email: claims.Sub.Email},
			)
			if err != nil {
				panic(err)
			}

			store.SetUser(ctx, u)

			ctx.Next()
		},
	)
}
