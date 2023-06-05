package persistence_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/token"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/auth/redis"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
	"github.com/redis/go-redis/v9"
)

type TokenSuite struct {
	suite.SuiteWithRedisConn
	Token func(client *redis.Client) repo.Token
}

func TestTokenSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(TokenSuite))
}

func (s *TokenSuite) SetupSuite() {
	s.Token = func(client *redis.Client) repo.Token {
		return persistence.NewRedisToken(client)
	}
}

func (s *TokenSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SaveTokenInput) (*token.Token, error)
		Input *repo.SaveTokenInput
		Ctx   context.Context
	}

	makeSut := func(client *redis.Client) *Sut {
		return &Sut{
			Sut: s.Token(client).Save,
			Input: &repo.SaveTokenInput{
				Token: fake.Token(),
			},
			Ctx: context.Background(),
		}
	}

	s.Run("should save the token", func(client *redis.Client) {
		sut := makeSut(client)

		token, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotNil(token)
		s.RedisExists(client, fmt.Sprintf("%s:%s", token.Email, token.Jti))
	})

	s.Run("should return an error if redis fails", func(client *redis.Client) {
		sut := makeSut(client)

		client.Close()

		token, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
		s.Nil(token)
	})
}

func (s *TokenSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeleteTokenInput) error
		Input *repo.DeleteTokenInput
		Ctx   context.Context
	}

	makeSut := func(client *redis.Client) *Sut {
		return &Sut{
			Sut: s.Token(client).Delete,
			Input: &repo.DeleteTokenInput{
				Email: userValue.Email(faker.Email()),
				Jti:   value.UUID(faker.UUIDHyphenated()),
			},
			Ctx: context.Background(),
		}
	}

	s.Run("should delete the token", func(client *redis.Client) {
		sut := makeSut(client)

		fixture.SaveToken(client, &fixture.TokenDeps{
			Token: &token.Token{
				Email: sut.Input.Email,
				Jti:   sut.Input.Jti,
			},
		})

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.RedisNotExists(client, fmt.Sprintf("%s:%s", sut.Input.Email, sut.Input.Jti))
	})

	s.Run("should return an error if key does not exist", func(client *redis.Client) {
		sut := makeSut(client)

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
		s.RedisNotExists(client, fmt.Sprintf("%s:%s", sut.Input.Email, sut.Input.Jti))
	})

	s.Run("should return an error if redis fails", func(client *redis.Client) {
		sut := makeSut(client)

		client.Close()

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *TokenSuite) TestDeleteAll() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeleteAllTokenInput) error
		Input *repo.DeleteAllTokenInput
		Ctx   context.Context
	}

	makeSut := func(client *redis.Client) *Sut {
		return &Sut{
			Sut: s.Token(client).DeleteAll,
			Input: &repo.DeleteAllTokenInput{
				Email: userValue.Email(faker.Email()),
			},
			Ctx: context.Background(),
		}
	}

	s.Run("should delete all tokens", func(client *redis.Client) {
		sut := makeSut(client)

		tokens := []*token.Token{
			fake.Token().SetEmail(sut.Input.Email),
			fake.Token().SetEmail(sut.Input.Email),
			fake.Token().SetEmail(sut.Input.Email),
		}

		for _, t := range tokens {
			fixture.SaveToken(client, &fixture.TokenDeps{
				Token: &token.Token{
					Email: t.Email,
					Jti:   t.Jti,
				},
			})
		}

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		for _, t := range tokens {
			s.RedisNotExists(client, fmt.Sprintf("%s:%s", sut.Input.Email, t.Jti))
		}
	})

	s.Run("should return an error if key does not exist", func(client *redis.Client) {
		sut := makeSut(client)

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})

	s.Run("should return an error if redis fails", func(client *redis.Client) {
		sut := makeSut(client)

		client.Close()

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *TokenSuite) TestFind() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.FindTokenInput) (*token.Token, error)
		Input *repo.FindTokenInput
		Ctx   context.Context
	}

	makeSut := func(client *redis.Client) *Sut {
		return &Sut{
			Sut: s.Token(client).Find,
			Input: &repo.FindTokenInput{
				Email: userValue.Email(faker.Email()),
				Jti:   value.UUID(faker.UUIDHyphenated()),
			},
			Ctx: context.Background(),
		}
	}

	s.Run("should find the token", func(client *redis.Client) {
		sut := makeSut(client)

		fixture.SaveToken(client, &fixture.TokenDeps{
			Token: &token.Token{
				Email: sut.Input.Email,
				Jti:   sut.Input.Jti,
			},
		})

		token, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotNil(token)
	})

	s.Run("should return an error if key does not exist", func(client *redis.Client) {
		sut := makeSut(client)

		token, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
		s.Nil(token)
	})

	s.Run("should return an error if redis fails", func(client *redis.Client) {
		sut := makeSut(client)

		client.Close()

		token, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
		s.Nil(token)
	})
}
