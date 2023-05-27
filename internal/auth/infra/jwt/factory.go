package jwt

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
	"github.com/christian-gama/nutrai-api/internal/core/infra/uuid"
)

func MakeAccessTokenGenerator() jwt.Generator {
	return NewGenerator(
		uuid.MakeGenerator(),
		jwt.AccessTokenType,
		env.Jwt.AccessExpire.Duration,
		persistence.MakeRedisToken(),
	)
}

func MakeRefreshTokenGenerator() jwt.Generator {
	return NewGenerator(
		uuid.MakeGenerator(),
		jwt.RefreshTokenType,
		env.Jwt.RefreshExpire.Duration,
		persistence.MakeRedisToken(),
	)
}

func MakeRefreshVerifier() jwt.Verifier {
	return NewVerifier(jwt.RefreshTokenType, persistence.MakeRedisToken())
}

func MakeAccessVerifier() jwt.Verifier {
	return NewVerifier(jwt.AccessTokenType, persistence.MakeRedisToken())
}
