package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/token"
	"github.com/redis/go-redis/v9"
)

type TokenDeps struct {
	Token *token.Token
}

func SaveToken(client *redis.Client, deps *TokenDeps) *TokenDeps {
	if deps == nil {
		deps = &TokenDeps{}
	}

	if deps.Token == nil {
		deps.Token = fake.Token()
	}

	token, err := persistence.NewRedisToken(client).
		Save(context.Background(), repo.SaveTokenInput{
			Token: deps.Token,
		})
	if err != nil {
		panic(fmt.Errorf("could not create token: %w", err))
	}

	deps.Token = token

	return deps
}
