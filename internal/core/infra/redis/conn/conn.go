package conn

import (
	"context"
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/retry"
	"github.com/redis/go-redis/v9"
)

type conn struct {
	client *redis.Client
}

func NewConn(db int) *conn {
	log.Loading("\tConnecting to Redis (%d)", db)

	client := redis.NewClient(&redis.Options{
		Addr:     addr(),
		Password: env.Redis.Password,
		DB:       db,
	})

	const attempts = 90
	err := retry.Retry(attempts, time.Second, func() error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.Ping(ctx).Result()
		return err
	})
	if err != nil {
		log.Fatalf("\tFailed to connect to redis after %d retries: %v", attempts, err)
	}

	return &conn{client}
}

func addr() string {
	return fmt.Sprintf("%s:%d", env.Redis.Host, env.Redis.Port)
}

func (c *conn) Close() error {
	c.check()
	return c.client.Close()
}

func (c *conn) check() {
	if c.client == nil {
		panic("redis connection is nil")
	}
}

func (c *conn) GetClient() *redis.Client {
	return c.client
}
