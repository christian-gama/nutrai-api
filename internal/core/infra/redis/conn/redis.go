package conn

import "github.com/redis/go-redis/v9"

var redisInstance *conn

// GetRedis returns the redis client.
func GetRedis() *redis.Client {
	return redisInstance.client
}
