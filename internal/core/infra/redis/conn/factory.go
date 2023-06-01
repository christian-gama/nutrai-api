package conn

import (
	"github.com/redis/go-redis/v9"
)

var defaultRedisConnection *redis.Conn

func MakeDefaultRedis() *redis.Conn {
	if defaultRedisConnection != nil {
		return defaultRedisConnection
	}

	defaultRedisConnection = NewConn(0)

	return defaultRedisConnection
}
