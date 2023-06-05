package asserts

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

// RedisExists asserts that the given key exists in the database.
func RedisExists(
	t *testing.T,
	client *redis.Client,
	key string,
	msgAndArgs ...any,
) bool {
	exists, err := client.Exists(context.Background(), key).Result()
	if err != nil {
		return assert.Fail(t, err.Error(), msgAndArgs...)
	}

	if msgAndArgs != nil {
		return assert.Equal(t, int64(1), exists, msgAndArgs...)
	}

	return assert.Equal(t, int64(1), exists, "expected key '%s' to exist", key)
}

// RedisNotExists asserts that the given key does not exist in the database.
func RedisNotExists(
	t *testing.T,
	client *redis.Client,
	key string,
	msgAndArgs ...any,
) bool {
	exists, err := client.Exists(context.Background(), key).Result()
	if err != nil {
		return assert.Fail(t, err.Error(), msgAndArgs...)
	}

	if msgAndArgs != nil {
		return assert.Equal(t, int64(0), exists, msgAndArgs...)
	}

	return assert.Equal(t, int64(0), exists, "expected key '%s' to not exist", key)
}

// RedisEqual asserts that the given key has the given value in the database.
func RedisEqual(
	t *testing.T,
	client *redis.Client,
	key string,
	expected interface{},
	msgAndArgs ...any,
) bool {
	actual, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return assert.Fail(t, err.Error(), msgAndArgs...)
	}

	if msgAndArgs != nil {
		return assert.Equal(t, expected, actual, msgAndArgs...)
	}

	return assert.Equal(t, expected, actual, "expected key '%s' to be equal to '%v'", key, expected)
}

// RedisNotEqual asserts that the given key does not have the given value in the database.
func RedisNotEqual(
	t *testing.T,
	client *redis.Client,
	key string,
	expected interface{},
	msgAndArgs ...any,
) bool {
	actual, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return assert.Fail(t, err.Error(), msgAndArgs...)
	}

	if msgAndArgs != nil {
		return assert.NotEqual(t, expected, actual, msgAndArgs...)
	}

	return assert.NotEqual(
		t,
		expected,
		actual,
		"expected key '%s' to not be equal to '%v'",
		key,
		expected,
	)
}
