package jwt_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	. "github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
	uuidMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/uuid"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type VerifierSuite struct {
	suite.Suite
}

func TestVerifierSuite(t *testing.T) {
	suite.RunUnitTest(t, new(VerifierSuite))
}

func (s *VerifierSuite) TestVerify() {
	type Mock struct {
		TokenRepo *repoMock.Token
	}

	type Sut struct {
		Sut            func(token value.Token, checkIsStored bool) (*jwt.Claims, error)
		TokenGenerator func(duration time.Duration) value.Token
		Data           *jwt.Subject
		Mock           *Mock
	}

	makeSut := func() *Sut {
		data := new(jwt.Subject)
		err := faker.FakeData(data)
		s.Require().NoError(err)

		mocks := &Mock{
			TokenRepo: repoMock.NewToken(s.T()),
		}

		tokenGenerator := func(duration time.Duration) value.Token {
			uuidMock := uuidMock.NewGenerator(s.T())
			uuidMock.On("Generate").Maybe().Return(coreValue.UUID("uuid"))

			mocks.TokenRepo.
				On("Save", mock.Anything, mock.Anything).
				Maybe().
				Return(token.NewToken(), nil)

			token, err := NewGenerator(
				uuidMock,
				jwt.AccessTokenType,
				duration,
				mocks.TokenRepo,
			).Generate(data, false)
			s.Require().NoError(err)

			return token
		}

		return &Sut{
			Sut:            NewVerifier(jwt.AccessTokenType, mocks.TokenRepo).Verify,
			TokenGenerator: tokenGenerator,
			Data:           data,
			Mock:           mocks,
		}
	}

	s.Run("should verify a valid JWT token succesfuly", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		claims, err := sut.Sut(token, false)
		s.NoError(err)

		s.Equal(jwt.AccessTokenType, claims.Type)
	})

	s.Run("should return an error if the token is invalid", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		claims, err := sut.Sut(token+"invalid", false)

		s.Error(err)
		s.Nil(claims)
	})

	s.Run("should return an error if the token is expired", func() {
		sut := makeSut()

		token := sut.TokenGenerator(-time.Hour)

		claims, err := sut.Sut(token, false)

		s.Require().Error(err)
		s.Contains(err.Error(), "expired")
		s.Nil(claims)
	})

	s.Run("should return an error if the token is not valid for the given type", func() {
		sut := makeSut()

		verifier := NewVerifier(jwt.RefreshTokenType, sut.Mock.TokenRepo)
		token := sut.TokenGenerator(time.Hour)

		claims, err := verifier.Verify(token, false)

		s.ErrorIs(err, ErrInvalidToken)
		s.Nil(claims)
	})

	s.Run("should call 'Find' method if 'checkIsStored' is true", func() {
		sut := makeSut()

		t := sut.TokenGenerator(time.Hour)

		sut.Mock.TokenRepo.On("Find", mock.Anything, mock.Anything).Return(token.NewToken(), nil)

		claims, err := sut.Sut(t, true)

		s.NoError(err)
		s.Equal(jwt.AccessTokenType, claims.Type)
	})

	s.Run("should return an error if the token is not stored", func() {
		sut := makeSut()

		t := sut.TokenGenerator(time.Hour)

		sut.Mock.TokenRepo.On("Find", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		claims, err := sut.Sut(t, true)

		s.ErrorIs(err, ErrInvalidToken)
		s.Nil(claims)
	})
}
