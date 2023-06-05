package persistence

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/redis/go-redis/v9"
)

type tokenImpl struct {
	redis *redis.Client
}

func NewRedisToken(redis *redis.Client) repo.Token {
	return &tokenImpl{
		redis: redis,
	}
}

// Delete implements repo.Token.
func (r *tokenImpl) Delete(ctx context.Context, input repo.DeleteTokenInput) error {
	count, err := r.redis.Del(ctx, fmt.Sprintf("%s:%s", input.Email.String(), input.Jti.String())).
		Result()
	if err != nil {
		return errors.InternalServerError("failed to delete token: %s", err.Error())
	}

	if count == 0 {
		return errors.NotFound("token not found")
	}

	return nil
}

// DeleteAll implements repo.Token.
func (r *tokenImpl) DeleteAll(ctx context.Context, input repo.DeleteAllTokenInput) error {
	pattern := fmt.Sprintf("%s:*", input.Email.String())

	keys, err := r.redis.Keys(ctx, pattern).Result()
	if err != nil {
		return errors.InternalServerError("failed to delete all tokens: %s", err.Error())
	}

	if len(keys) == 0 {
		return errors.NotFound("tokens not found")
	}

	return r.redis.Del(ctx, keys...).Err()
}

// Find implements repo.Token.
func (r *tokenImpl) Find(
	ctx context.Context,
	input repo.FindTokenInput,
) (*token.Token, error) {
	value, err := r.redis.Get(ctx, fmt.Sprintf("%s:%s", input.Email.String(), input.Jti.String())).
		Result()
	if err != nil {
		return nil, errors.InternalServerError("failed to find token: %s", err.Error())
	}

	if value == "" {
		return nil, errors.NotFound("token not found")
	}

	expiresAt, err := r.redis.TTL(ctx, input.Email.String()).Result()
	if err != nil {
		return nil, errors.InternalServerError("failed to find token: %s", err.Error())
	}

	return token.NewToken().
		SetEmail(input.Email).
		SetJti(input.Jti).
		SetExpiresAt(expiresAt), nil
}

// Save implements repo.Token.
func (r *tokenImpl) Save(
	ctx context.Context,
	input repo.SaveTokenInput,
) (*token.Token, error) {
	key := fmt.Sprintf("%s:%s", input.Token.Email.String(), input.Token.Jti.String())
	if err := r.redis.Set(ctx, key, true, input.Token.ExpiresAt).Err(); err != nil {
		return nil, err
	}

	return input.Token, nil
}
