package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
	userMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
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
	type Mocks struct {
		UserRepo *userMock.User
	}

	type Sut struct {
		Sut   command.DeleteUserHandler
		Ctx   context.Context
		Input *command.DeleteUserInput
		Mocks *Mocks
	}

	makeSut := func() *Sut {
		userRepo := userMock.NewUser(s.T())
		input := fake.DeleteUserInput()
		sut := command.NewDeleteUserHandler(userRepo)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mocks: &Mocks{
				UserRepo: userRepo,
			},
		}
	}

	s.Run("Should delete the user if succeed", func() {
		sut := makeSut()

		user := userFake.User()
		sut.Mocks.UserRepo.On("Find", sut.Ctx, mock.Anything).Return(user, nil)
		sut.Mocks.UserRepo.On("Delete", sut.Ctx, mock.Anything).Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the user does not exist", func() {
		sut := makeSut()

		sut.Mocks.UserRepo.On("Find", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
