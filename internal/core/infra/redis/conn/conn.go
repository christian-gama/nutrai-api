package conn

import (
	"context"
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/pkg/retry"
	"github.com/redis/go-redis/v9"
)

type conn struct {
	redis *redis.Conn
	log   logger.Logger
}

func NewConn(log logger.Logger) *conn {
	return &conn{
		log: log,
	}
}

func (c *conn) Open(db int) *redis.Conn {
	c.log.Loading("\tConnecting to Redis (%d)", db)

	redis := redis.NewClient(&redis.Options{
		Addr:     c.addr(),
		Password: env.Redis.Password,
		DB:       db,
	})

	const attempts = 90
	err := retry.Retry(attempts, time.Second, func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		c.redis = redis.Conn()
		_, err := c.redis.Ping(ctx).Result()
		return err
	})
	if err != nil {
		c.log.Fatalf("\tFailed to connect to redis after %d retries: %v", attempts, err)
	}

	return c.redis
}

func (c *conn) addr() string {
	return fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port)
}
