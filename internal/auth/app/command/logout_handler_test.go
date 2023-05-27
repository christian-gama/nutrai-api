package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
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
		}

		input := fake.LogoutInput()

		sut := command.NewLogoutHandler(mock.TokenRepo)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should delete the token if succeed", func() {
		sut := makeSut()

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(token.NewToken(), nil)

		sut.Mock.TokenRepo.
			On("Delete", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the user does not exist", func() {
		sut := makeSut()

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(token.NewToken(), nil)

		sut.Mock.TokenRepo.
			On("Delete", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return an error if the token does not exist", func() {
		sut := makeSut()

		sut.Mock.TokenRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
