package service_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/service"
	jwtFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/jwt"
	jwtMocks "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RefreshTokenHandlerSuite struct {
	suite.Suite
}

func TestRefreshTokenHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RefreshTokenHandlerSuite))
}

func (s *RefreshTokenHandlerSuite) TestHandle() {
	type Mocks struct {
		AccessToken *jwtMocks.Generator
		Verifier    *jwtMocks.Verifier
	}

	type Sut struct {
		Sut   func(ctx context.Context, input *service.RefreshTokenInput) (*service.RefreshTokenOutput, error)
		Mocks *Mocks
		Input *service.RefreshTokenInput
	}

	makeSut := func() *Sut {
		mocks := &Mocks{
			AccessToken: jwtMocks.NewGenerator(s.T()),
			Verifier:    jwtMocks.NewVerifier(s.T()),
		}

		sut := service.NewRefreshTokenHandler(
			mocks.AccessToken,
			mocks.Verifier,
		)

		input := fake.RefreshTokenInput()

		return &Sut{
			Sut:   sut.Handle,
			Mocks: mocks,
			Input: input,
		}
	}

	s.Run("should return an access and refresh token on success", func() {
		sut := makeSut()

		sut.Mocks.Verifier.
			On("Verify", mock.Anything).
			Return(jwtFake.AccessTokenClaims(), nil)

		accessToken := value.Token("access")
		sut.Mocks.AccessToken.
			On("Generate", mock.Anything).
			Return(accessToken, nil)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.Nil(err)
		s.EqualValues(accessToken, output.Access)
	})

	s.Run("should return an error if the refresh token is invalid", func() {
		sut := makeSut()

		sut.Mocks.Verifier.
			On("Verify", mock.Anything).
			Return(nil, assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(output)
	})

	s.Run("should return an error if the access token can't be generated", func() {
		sut := makeSut()

		sut.Mocks.Verifier.
			On("Verify", mock.Anything).
			Return(jwtFake.AccessTokenClaims(), nil)

		sut.Mocks.AccessToken.
			On("Generate", mock.Anything).
			Return(value.Token(""), assert.AnError)

		output, err := sut.Sut(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(output)
	})
}
