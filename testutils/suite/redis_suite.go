package suite

import (
	"sync"

	"github.com/christian-gama/nutrai-api/internal/core/infra/redis/conn"
	"github.com/christian-gama/nutrai-api/testutils/suite/asserts"
	"github.com/redis/go-redis/v9"
)

// SuiteWithRedisConn is the base suite for all test suites. It provides helper methods for
// testing.
// It provides helper methods for testing and wraps each test with a redis connection, setting
// up the connection before the test and closing it after the test is done. Every database index is
// incremented before the test and decremented after the test is done.
type SuiteWithRedisConn struct {
	Suite
}

// Run runs a test with a redis connection, setting up the connection before the test and closing
// it after the test is done. Every database index is incremented before the test and decremented
// after the test is done.
func (s *SuiteWithRedisConn) Run(
	name string,
	f func(client *redis.Client),
) bool {
	return s.Suite.Run(name, func() {
		redisDBManagerInstance.increase()

		client := conn.NewConn(redisDBManagerInstance.db).Client()
		defer func() {
			client.Close()
			redisDBManagerInstance.decrease()
		}()

		f(client)
	})
}

func (s *SuiteWithRedisConn) RedisExists(
	client *redis.Client,
	key string,
	msgAndArgs ...any,
) bool {
	return asserts.RedisExists(s.T(), client, key, msgAndArgs...)
}

func (s *SuiteWithRedisConn) RedisNotExists(
	client *redis.Client,
	key string,
	msgAndArgs ...any,
) bool {
	return asserts.RedisNotExists(s.T(), client, key, msgAndArgs...)
}

func (s *SuiteWithRedisConn) RedisEqual(
	client *redis.Client,
	key string,
	expected interface{},
	msgAndArgs ...any,
) bool {
	return asserts.RedisEqual(s.T(), client, key, expected, msgAndArgs...)
}

func (s *SuiteWithRedisConn) RedisNotEqual(
	client *redis.Client,
	key string,
	expected interface{},
	msgAndArgs ...any,
) bool {
	return asserts.RedisNotEqual(s.T(), client, key, expected, msgAndArgs...)
}

// redisDBManager is a manager for the redis database index.
type redisDBManager struct {
	db int
	mu sync.Mutex
}

// increase increments the database index by 1, locking the mutex before and unlocking it after.
func (r *redisDBManager) increase() {
	r.mu.Lock()
	r.db++
	r.mu.Unlock()
}

// decrease decrements the database index by 1, locking the mutex before and unlocking it after.
func (r *redisDBManager) decrease() {
	r.mu.Lock()
	r.db--
	r.mu.Unlock()
}

var redisDBManagerInstance = &redisDBManager{
	db: 0,
	mu: sync.Mutex{},
}
