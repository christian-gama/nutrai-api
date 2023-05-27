package jwt_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	. "github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/uuid"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
)

type VerifierSuite struct {
	suite.Suite
}

func TestVerifierSuite(t *testing.T) {
	suite.RunUnitTest(t, new(VerifierSuite))
}

func (s *VerifierSuite) TestGenerate() {
	type Sut struct {
		Sut            func(token value.Token) (*jwt.Claims, error)
		TokenGenerator func(duration time.Duration) value.Token
		Data           *jwt.Subject
	}

	makeSut := func() *Sut {
		data := new(jwt.Subject)
		err := faker.FakeData(data)
		s.Require().NoError(err)

		tokenGenerator := func(duration time.Duration) value.Token {
			uuidMock := mocks.NewGenerator(s.T())
			uuidMock.On("Generate").Maybe().Return(coreValue.UUID("uuid"))

			token, err := NewGenerator(uuidMock, jwt.AccessTokenType, duration).Generate(data)
			s.Require().NoError(err)

			return token
		}

		return &Sut{
			Sut:            NewVerifier(jwt.AccessTokenType).Verify,
			TokenGenerator: tokenGenerator,
		}
	}

	s.Run("should verify a valid JWT token succesfuly", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		claims, err := sut.Sut(token)
		s.NoError(err)

		s.Equal(jwt.AccessTokenType, claims.Type)
	})

	s.Run("should return an error if the token is invalid", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		claims, err := sut.Sut(token + "invalid")
		s.Error(err)
		s.Nil(claims)
	})

	s.Run("should return an error if the token is expired", func() {
		sut := makeSut()

		token := sut.TokenGenerator(-time.Hour)

		claims, err := sut.Sut(token)
		s.Require().Error(err)
		s.Contains(err.Error(), "expired")
		s.Nil(claims)
	})

	s.Run("should return an error if the token is not valid for the given audience", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)
		env.App.Host = "invalid"

		claims, err := sut.Sut(token)
		s.Require().Error(err)
		s.Nil(claims)
	})

	s.Run("should return an error if the token is not valid for the given type", func() {
		sut := makeSut()

		verifier := NewVerifier(jwt.RefreshTokenType)
		token := sut.TokenGenerator(time.Hour)

		claims, err := verifier.Verify(token)
		s.ErrorIs(err, ErrInvalidToken)
		s.Nil(claims)
	})
}
