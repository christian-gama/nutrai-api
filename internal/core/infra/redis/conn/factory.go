package conn

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/redis/go-redis/v9"
)

var defaultRedisConnection *redis.Conn

func MakeDefaultRedis() *redis.Conn {
	if defaultRedisConnection != nil {
		return defaultRedisConnection
	}

	log := log.MakeWithCaller()
	conn := NewConn(log)
	defaultRedisConnection = conn.Open(0)

	return defaultRedisConnection
}
