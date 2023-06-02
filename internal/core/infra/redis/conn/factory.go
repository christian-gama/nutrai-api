package conn

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func MakeRedis() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	redisClient = NewConn(0)

	return redisClient
}
