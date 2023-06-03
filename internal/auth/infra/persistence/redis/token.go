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
	if input.Email == "" || input.Jti == "" {
		return errors.InternalServerError("email and jti are required to delete a token")
	}

	return r.redis.Del(ctx, fmt.Sprintf("%s:%s", input.Email.String(), input.Jti.String())).Err()
}

// DeleteAll implements repo.Token.
func (r *tokenImpl) DeleteAll(ctx context.Context, input repo.DeleteAllTokenInput) error {
	if input.Email == "" {
		return errors.InternalServerError("email is required to delete all tokens")
	}

	pattern := fmt.Sprintf("%s:*", input.Email.String())

	keysCmd := r.redis.Keys(ctx, pattern)
	if keysCmd.Err() != nil {
		return errors.InternalServerError(keysCmd.Err().Error())
	}

	if len(keysCmd.Val()) == 0 {
		return nil
	}

	return r.redis.Del(ctx, keysCmd.Val()...).Err()
}

// Find implements repo.Token.
func (r *tokenImpl) Find(
	ctx context.Context,
	input repo.FindTokenInput,
) (*token.Token, error) {
	if input.Email == "" || input.Jti == "" {
		return nil, errors.InternalServerError("email and jti are required to find a token")
	}

	strCmd := r.redis.Get(ctx, fmt.Sprintf("%s:%s", input.Email.String(), input.Jti.String()))
	if strCmd.Err() != nil {
		return nil, errors.Unauthorized("invalid token")
	}

	ttlCmd := r.redis.TTL(ctx, input.Email.String())
	if ttlCmd.Err() != nil {
		return nil, errors.Unauthorized("invalid token")
	}

	return token.NewToken().
		SetEmail(input.Email).
		SetJti(input.Jti).
		SetExpiresAt(ttlCmd.Val()), nil
}

// Save implements repo.Token.
func (r *tokenImpl) Save(
	ctx context.Context,
	input repo.SaveTokenInput,
) (*token.Token, error) {
	if err := r.redis.Set(
		ctx,
		fmt.Sprintf("%s:%s", input.Token.Email.String(), input.Token.Jti.String()),
		true,
		input.Token.ExpiresAt,
	).Err(); err != nil {
		return nil, err
	}

	return input.Token, nil
}
