package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	userMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteUserSuite struct {
	suite.Suite
}

func TestDeleteUserSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteUserSuite))
}

func (s *DeleteUserSuite) TestDeleteUser() {
	type Mock struct {
		UserRepo *userMock.User
	}

	type Sut struct {
		Sut   command.DeleteMeHandler
		Ctx   context.Context
		Input *command.DeleteMeInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			UserRepo: userMock.NewUser(s.T()),
		}

		input := fake.DeleteMeInput()

		sut := command.NewDeleteMeHandler(mock.UserRepo)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should delete the user if succeed", func() {
		sut := makeSut()

		user := userFake.User()
		sut.Mock.UserRepo.On("Find", sut.Ctx, mock.Anything).Return(user, nil)
		sut.Mock.UserRepo.On("Delete", sut.Ctx, mock.Anything).Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the user does not exist", func() {
		sut := makeSut()

		sut.Mock.UserRepo.On("Find", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
