package jwt_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
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
		Sut            func(token value.Token) (*jwt.Payload, error)
		TokenGenerator func(duration time.Duration) value.Token
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
			Sut:            NewVerifier().Verify,
			TokenGenerator: tokenGenerator,
		}
	}

	s.Run("should verify a valid JWT token succesfuly", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		payload, err := sut.Sut(token)
		s.NoError(err)

		s.Equal(jwt.AccessTokenType, payload.Type)
	})

	s.Run("should return an error if the token is invalid", func() {
		sut := makeSut()

		token := sut.TokenGenerator(time.Hour)

		payload, err := sut.Sut(token + "invalid")
		s.Error(err)
		s.Nil(payload)
	})

	s.Run("should return an error if the token is expired", func() {
		sut := makeSut()

		token := sut.TokenGenerator(-time.Hour)

		payload, err := sut.Sut(token)
		s.Require().Error(err)
		s.Contains(err.Error(), "expired")
		s.Nil(payload)
	})
}
