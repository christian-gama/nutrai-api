package conn

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/redis/go-redis/v9"
)

var redisInstance *conn

// GetRedis returns the redis client.
func GetRedis() *redis.Client {
	if redisInstance == nil {
		log.Fatal("redis connection does not exist - did you forget to initialize it?")
	}

	return redisInstance.Client()
}
