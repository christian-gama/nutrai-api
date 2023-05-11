package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
	hasherMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/hasher"
	userMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CheckCredentialsSuite struct {
	suite.Suite
}

func TestCheckCredentialsSuite(t *testing.T) {
	suite.RunUnitTest(t, new(CheckCredentialsSuite))
}

func (s *CheckCredentialsSuite) TestCheckCredentials() {
	type Mocks struct {
		Hasher   *hasherMock.Hasher
		UserRepo *userMock.User
	}

	type Sut struct {
		Sut   command.CheckCredentialsHandler
		Ctx   context.Context
		Input *command.CheckCredentialsInput
		Mocks *Mocks
	}

	makeSut := func() *Sut {
		hasher := hasherMock.NewHasher(s.T())
		userRepo := userMock.NewUser(s.T())
		input := fake.CheckCredentialsInput()
		sut := command.NewCheckCredentialsHandler(userRepo, hasher)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mocks: &Mocks{
				Hasher:   hasher,
				UserRepo: userRepo,
			},
		}
	}

	s.Run("Should return the user if succeed", func() {
		sut := makeSut()

		user := userFake.User()
		sut.Mocks.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(user, nil)
		sut.Mocks.Hasher.On("Compare", value.Password(sut.Input.Password), user.Password).Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the user does not exist", func() {
		sut := makeSut()

		sut.Mocks.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return an error if the password does not match", func() {
		sut := makeSut()

		user := userFake.User()
		sut.Mocks.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(user, nil)
		sut.Mocks.Hasher.On("Compare", value.Password(sut.Input.Password), user.Password).Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsInvalid(err)
	})

	s.Run("Should return an error if the password is empty", func() {
		sut := makeSut()

		user := userFake.User()
		sut.Mocks.UserRepo.On("FindByEmail", sut.Ctx, mock.Anything).Return(user, nil)
		sut.Input.Password = ""

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})
}
