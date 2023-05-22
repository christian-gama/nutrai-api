package jwt

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/infra/uuid"
)

func MakeAccessTokenGenerator() jwt.Generator {
	return NewGenerator(uuid.MakeGenerator(), jwt.AccessTokenType, env.Jwt.AccessExpire.Duration)
}

func MakeRefreshTokenGenerator() jwt.Generator {
	return NewGenerator(uuid.MakeGenerator(), jwt.RefreshTokenType, env.Jwt.RefreshExpire.Duration)
}

func MakeVerifier() jwt.Verifier {
	return NewVerifier()
}
