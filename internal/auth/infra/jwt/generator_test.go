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

type GeneratorSuite struct {
	suite.Suite
}

func TestGeneratorSuite(t *testing.T) {
	suite.RunUnitTest(t, new(GeneratorSuite))
}

func (s *GeneratorSuite) TestGenerate() {
	type Sut struct {
		Sut  func(data *jwt.Subject) (value.Token, error)
		Data *jwt.Subject
	}

	makeSut := func() *Sut {
		data := new(jwt.Subject)
		err := faker.FakeData(data)
		s.Require().NoError(err)

		uuidMock := mocks.NewGenerator(s.T())
		uuidMock.On("Generate").Maybe().Return(coreValue.UUID("uuid"))

		return &Sut{
			Sut: NewGenerator(
				uuidMock,
				jwt.AccessTokenType,
				time.Hour,
			).Generate,
			Data: data,
		}
	}

	s.Run("should generate a new JWT token", func() {
		sut := makeSut()

		token, err := sut.Sut(sut.Data)
		s.NoError(err)
		s.NotEmpty(token)
		s.Regexp(`^([a-zA-Z0-9_-]+\.){2}[a-zA-Z0-9_-]+$`, token)
	})

	s.Run("should return an error if the subject is nil", func() {
		sut := makeSut()

		token, err := sut.Sut(nil)
		s.Error(err)
		s.Empty(token)
	})

	s.Run("should return an error if the subject's email is invalid", func() {
		sut := makeSut()

		sut.Data.Email = coreValue.Email("invalid")

		token, err := sut.Sut(sut.Data)
		s.Error(err)
		s.Empty(token)
	})
}
