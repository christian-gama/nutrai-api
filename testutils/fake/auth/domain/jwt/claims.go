package fake

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/go-faker/faker/v4"
)

func tokenClaims() *jwt.Claims {
	return &jwt.Claims{
		Aud: "aud",
		Exp: time.Now().Add(time.Hour).Unix(),
		Iat: time.Now().Unix(),
		Iss: "iss",
		Jti: coreValue.UUID(faker.UUIDHyphenated()),
		Nbf: time.Now().Unix(),
		Sub: jwt.Subject{
			Email: value.Email(faker.Email()),
		},
		Type: "refresh",
	}
}

func AccessTokenClaims() *jwt.Claims {
	token := tokenClaims()
	token.Type = "access"

	return token
}

func RefreshTokenClaims() *jwt.Claims {
	token := tokenClaims()
	token.Type = "refresh"

	return token
}
