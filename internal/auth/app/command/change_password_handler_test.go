package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	hasherMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/hasher"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
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
		Hasher   *hasherMock.Hasher
		UserRepo *userRepoMock.User
	}

	type Sut struct {
		Sut   command.ChangePasswordHandler
		Ctx   context.Context
		Input *command.ChangePasswordInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			Hasher:   hasherMock.NewHasher(s.T()),
			UserRepo: userRepoMock.NewUser(s.T()),
		}

		input := fake.ChangePasswordInput()

		sut := command.NewChangePasswordHandler(mock.UserRepo, mock.Hasher)

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

		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should call hasher.Hash with the password", func() {
		sut := makeSut()

		password := sut.Input.Password
		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.Hasher.AssertCalled(s.T(), "Hash", password)
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

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password(""), assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password(""), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when updating the password fails", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should call userRepo.Update with the new hashed password", func() {
		sut := makeSut()

		sut.Mock.UserRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(userFake.User(), nil)

		hashedPassword := value.Password("hashed")
		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(hashedPassword, nil)

		sut.Mock.UserRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.UserRepo.AssertCalled(
			s.T(),
			"Update",
			sut.Ctx,
			mock.MatchedBy(func(input repo.UpdateUserInput) bool {
				return input.User.Password == hashedPassword
			}),
		)
	})
}
