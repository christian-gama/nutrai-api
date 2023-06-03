package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LogoutAllSuite struct {
	suite.Suite
}

func TestLogoutAllSuite(t *testing.T) {
	suite.RunUnitTest(t, new(LogoutAllSuite))
}

func (s *LogoutAllSuite) TestLogoutAll() {
	type Mock struct {
		TokenRepo *repoMock.Token
	}

	type Sut struct {
		Sut   command.LogoutAllHandler
		Ctx   context.Context
		Input *command.LogoutAllInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			TokenRepo: repoMock.NewToken(s.T()),
		}

		input := fake.LogoutAllInput()

		sut := command.NewLogoutAllHandler(mock.TokenRepo)

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
			On("DeleteAll", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the user does not exist", func() {
		sut := makeSut()

		sut.Mock.TokenRepo.
			On("DeleteAll", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsInternalServerError(err)
	})

	s.Run("Should return an error if deleting all tokens fails", func() {
		sut := makeSut()

		sut.Mock.TokenRepo.
			On("DeleteAll", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsInternalServerError(err)
	})
}
