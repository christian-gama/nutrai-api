package service_test

import (
	"context"
	"testing"

	userCmd "github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/service"
	jwtMocks "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/jwt"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RegisterHandlerSuite struct {
	suite.Suite
}

func TestRegisterHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RegisterHandlerSuite))
}

func (s *RegisterHandlerSuite) TestHandle() {
	type Mocks struct {
		AccessToken     *jwtMocks.Generator
		RefreshToken    *jwtMocks.Generator
		SaveUserHandler *cmdMock.Handler[*userCmd.SaveUserInput]
	}

	type Sut struct {
		Sut   func(ctx context.Context, input *service.RegisterInput) (*service.RegisterOutput, error)
		Mocks *Mocks
		Input *service.RegisterInput
	}

	makeSut := func() *Sut {
		mocks := &Mocks{
			AccessToken:     jwtMocks.NewGenerator(s.T()),
			RefreshToken:    jwtMocks.NewGenerator(s.T()),
			SaveUserHandler: cmdMock.NewHandler[*userCmd.SaveUserInput](s.T()),
		}

		sut := service.NewRegisterHandler(
			mocks.AccessToken,
			mocks.RefreshToken,
			mocks.SaveUserHandler,
		)

		input := fake.RegisterInput()

		return &Sut{
			Sut:   sut.Handle,
			Mocks: mocks,
			Input: input,
		}
	}

	s.Run("should return an access and refresh token on success", func() {
		sut := makeSut()

		sut.Mocks.SaveUserHandler.
			On("Handle", context.Background(), mock.Anything).
			Return(nil)

		accessToken := value.Token("access")
		sut.Mocks.AccessToken.
			On("Generate", mock.Anything).
			Return(accessToken, nil)

		refreshToken := value.Token("refresh")
		sut.Mocks.RefreshToken.
			On("Generate", mock.Anything).
			Return(refreshToken, nil)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(err)
		s.EqualValues(accessToken, output.Access)
		s.EqualValues(refreshToken, output.Refresh)
	})

	s.Run("should return an error if check credentials handler returns an error", func() {
		sut := makeSut()

		sut.Mocks.SaveUserHandler.
			On("Handle", context.Background(), mock.Anything).
			Return(assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(output)
		s.ErrorIs(err, assert.AnError)
	})

	s.Run("should return an error if access token generator returns an error", func() {
		sut := makeSut()

		sut.Mocks.SaveUserHandler.
			On("Handle", context.Background(), mock.Anything).
			Return(nil)

		sut.Mocks.AccessToken.
			On("Generate", mock.Anything).
			Return(value.Token(""), assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(output)
		s.ErrorIs(err, assert.AnError)
	})

	s.Run("should return an error if refresh token generator returns an error", func() {
		sut := makeSut()

		sut.Mocks.SaveUserHandler.
			On("Handle", context.Background(), mock.Anything).
			Return(nil)

		sut.Mocks.AccessToken.
			On("Generate", mock.Anything).
			Return(value.Token("access"), nil)

		sut.Mocks.RefreshToken.
			On("Generate", mock.Anything).
			Return(value.Token(""), assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(output)
		s.ErrorIs(err, assert.AnError)
	})
}
