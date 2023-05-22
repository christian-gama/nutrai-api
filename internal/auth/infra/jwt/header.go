package jwt

import (
	"strings"

	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/gin-gonic/gin"
)

// GetTokenFromHeader gets the token from the Authorization header.
func GetTokenFromHeader(ctx *gin.Context) (value.Token, error) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		return "", ErrMissingAuthorizationHeader
	}

	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return "", ErrInvalidAuthorizationHeader
	}

	return value.Token(token[1]), nil
}
