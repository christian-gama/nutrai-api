package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	hasherMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/hasher"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SaveUserHandlerSuite struct {
	suite.Suite
}

func TestSaveUserHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SaveUserHandlerSuite))
}

func (s *SaveUserHandlerSuite) TestSaveHandler() {
	type Mock struct {
		Hasher   *hasherMock.Hasher
		UserRepo *userRepoMock.User
	}

	type Sut struct {
		Sut   command.SaveUserHandler
		Ctx   context.Context
		Input *command.SaveUserInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			Hasher:   hasherMock.NewHasher(s.T()),
			UserRepo: userRepoMock.NewUser(s.T()),
		}

		input := fake.SaveUserInput()

		sut := command.NewSaveUserHandler(mock.UserRepo, mock.Hasher)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return nil when save user successfully", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(user.NewUser(), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should call hasher.Hash with the password", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", sut.Input.Password).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(user.NewUser(), nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.Hasher.AssertCalled(s.T(), "Hash", sut.Input.Password)
	})

	s.Run("Should return error when hashing password fails", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password(""), assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password(""), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when save the user fails", func() {
		sut := makeSut()

		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(value.Password("hashed"), nil)

		sut.Mock.UserRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should call userRepo.Save with the new hashed password", func() {
		sut := makeSut()

		hashedPassword := value.Password("hashed")
		sut.Mock.Hasher.
			On("Hash", mock.Anything).
			Return(hashedPassword, nil)

		sut.Mock.UserRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(user.NewUser(), nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.UserRepo.AssertCalled(
			s.T(),
			"Save",
			sut.Ctx,
			mock.MatchedBy(func(input repo.SaveUserInput) bool {
				return input.User.Password == hashedPassword
			}),
		)
	})
}
