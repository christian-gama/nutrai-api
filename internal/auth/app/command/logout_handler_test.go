package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	jwtMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/jwt"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LogoutSuite struct {
	suite.Suite
}

func TestLogoutSuite(t *testing.T) {
	suite.RunUnitTest(t, new(LogoutSuite))
}

func (s *LogoutSuite) TestLogout() {
	type Mock struct {
		TokenRepo *repoMock.Token
		Verifier  *jwtMock.Verifier
	}

	type Sut struct {
		Sut   command.LogoutHandler
		Ctx   context.Context
		Input *command.LogoutInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			TokenRepo: repoMock.NewToken(s.T()),
			Verifier:  jwtMock.NewVerifier(s.T()),
		}

		input := fake.LogoutInput()

		sut := command.NewLogoutHandler(mock.TokenRepo, mock.Verifier)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should delete the token if succeed", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Refresh, true).
			Return(&jwt.Claims{
				Sub: jwt.Subject{
					Email: sut.Input.User.Email,
				},
				Jti: value.UUID(faker.UUIDHyphenated()),
			}, nil)

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(token.NewToken(), nil)

		sut.Mock.TokenRepo.
			On("Delete", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if when trying to delete the token fails", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Refresh, true).
			Return(&jwt.Claims{
				Sub: jwt.Subject{
					Email: sut.Input.User.Email,
				},
				Jti: value.UUID(faker.UUIDHyphenated()),
			}, nil)

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(token.NewToken(), nil)

		sut.Mock.TokenRepo.
			On("Delete", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsInternalServerError(err)
	})

	s.Run("Should return an error if the token does not exist", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Refresh, true).
			Return(&jwt.Claims{
				Sub: jwt.Subject{
					Email: sut.Input.User.Email,
				},
				Jti: value.UUID(faker.UUIDHyphenated()),
			}, nil)

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsUnauthorized(err)
	})

	s.Run("Should return an error if the verifier returns an error", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Refresh, true).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsUnauthorized(err)
	})
}
