package jwt

import (
	"strings"

	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	"github.com/gin-gonic/gin"
)

// GetTokenFromHeader is a function that extracts the JWT token from the Authorization header
// of an incoming HTTP request. It first gets the bearer token from the header and then retrieves
// the actual JWT token from it. If either operation fails, an error is returned.
func GetTokenFromHeader(ctx *gin.Context) (value.Token, error) {
	bearerToken, err := getBearerToken(ctx)
	if err != nil {
		return "", err
	}

	return getToken(bearerToken)
}

// getBearerToken is a helper function that extracts the 'Bearer' token from the Authorization
// header. This token is a type of HTTP authentication scheme that involves security tokens called
// bearer tokens.
// If the Authorization header is missing or empty, it returns an error.
func getBearerToken(ctx *gin.Context) (string, error) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		return "", ErrMissingAuthorizationHeader
	}

	return bearerToken, nil
}

// getToken is a helper function that parses the bearer token to extract the actual JWT token.
// In the 'Bearer' scheme, the HTTP Authorization header has two parts: 'Bearer' and the 'token'.
// This function splits the bearer token by space, checks its validity, and returns the second part
// which is the actual JWT token. If the bearer token is invalid (it doesn't have exactly two
// parts), an error is returned.
func getToken(bearerToken string) (value.Token, error) {
	token := strings.Split(bearerToken, " ")
	if len(token) != 2 {
		return "", ErrInvalidAuthorizationHeader
	}

	return value.Token(token[1]), nil
}
