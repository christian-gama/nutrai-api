package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
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
	return r.redis.Del(ctx, input.Email.String()).Err()
}

// Find implements repo.Token.
func (r *tokenImpl) Find(
	ctx context.Context,
	input repo.FindTokenInput,
) (*token.Token, error) {
	strCmd := r.redis.Get(ctx, input.Email.String())
	if strCmd.Err() != nil {
		return nil, errors.Unauthorized("invalid token")
	}

	ttlCmd := r.redis.TTL(ctx, input.Email.String())
	if ttlCmd.Err() != nil {
		return nil, errors.Unauthorized("invalid token")
	}

	return token.NewToken().
		SetEmail(input.Email).
		SetJti(coreValue.UUID(strCmd.Val())).
		SetExpiresAt(ttlCmd.Val()), nil
}

// Save implements repo.Token.
func (r *tokenImpl) Save(
	ctx context.Context,
	input repo.SaveTokenInput,
) (*token.Token, error) {
	if err := r.redis.Set(
		ctx,
		input.Token.Email.String(),
		input.Token.Jti.String(),
		input.Token.ExpiresAt,
	).Err(); err != nil {
		return nil, err
	}

	return input.Token, nil
}
