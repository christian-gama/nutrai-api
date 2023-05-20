package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
	serviceMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/service"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ChangePasswordHandlerSuite struct {
	suite.Suite
}

func TestChangePasswordHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ChangePasswordHandlerSuite))
}

func (s *ChangePasswordHandlerSuite) TestSaveHandler() {
	type Mock struct {
		HashPasswordHandler *serviceMock.Handler[*service.HashPasswordInput, *service.HashPasswordOutput]
		UserRepo            *userRepoMock.User
	}

	type Sut struct {
		Sut   command.ChangePasswordHandler
		Ctx   context.Context
		Input *command.ChangePasswordInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			HashPasswordHandler: serviceMock.NewHandler[*service.HashPasswordInput, *service.HashPasswordOutput](
				s.T(),
			),
			UserRepo: userRepoMock.NewUser(s.T()),
		}

		input := fake.ChangePasswordInput()

		sut := command.NewChangePasswordHandler(mock.UserRepo, mock.HashPasswordHandler)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return nil when change password successfully", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should call hashPasswordHandler.Handle with the password", func() {
		sut := makeSut()

		password := sut.Input.Password
		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.HashPasswordHandler.AssertCalled(
			s.T(),
			"Handle",
			sut.Ctx,
			&service.HashPasswordInput{Password: password},
		)
	})

	s.Run("Should return error when user not found", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when hashing password fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			// Must return a pointer to output because of generics.
			Return(&service.HashPasswordOutput{}, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: ""}, nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when updating the password fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
