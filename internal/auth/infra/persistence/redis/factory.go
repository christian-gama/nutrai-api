package persistence

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
)

func MakeRedisToken() repo.Token {
	return NewRedisToken(conn.MakeRedis())
}
