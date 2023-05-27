package jwt_test

import (
	"testing"
	"time"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	. "github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	repoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
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
	type Mock struct {
		TokenRepo *repoMock.Token
	}

	type Sut struct {
		Sut  func(data *jwt.Subject, persist bool) (jwtValue.Token, error)
		Data *jwt.Subject
		Mock *Mock
	}

	makeSut := func() *Sut {
		data := new(jwt.Subject)
		err := faker.FakeData(data)
		s.Require().NoError(err)

		mock := &Mock{
			TokenRepo: repoMock.NewToken(s.T()),
		}

		uuidMock := mocks.NewGenerator(s.T())
		uuidMock.On("Generate").Maybe().Return(coreValue.UUID("uuid"))

		return &Sut{
			Sut: NewGenerator(
				uuidMock,
				jwt.AccessTokenType,
				time.Hour,
				mock.TokenRepo,
			).Generate,
			Data: data,
		}
	}

	s.Run("should generate a new JWT token", func() {
		sut := makeSut()

		token, err := sut.Sut(sut.Data, false)

		s.NoError(err)
		s.NotEmpty(token)
		s.Regexp(`^([a-zA-Z0-9_-]+\.){2}[a-zA-Z0-9_-]+$`, token)
	})

	s.Run("should return an error if the subject is nil", func() {
		sut := makeSut()

		token, err := sut.Sut(nil, false)

		s.Error(err)
		s.Empty(token)
	})

	s.Run("should return an error if the subject's email is empty", func() {
		sut := makeSut()

		sut.Data.Email = userValue.Email("")

		token, err := sut.Sut(sut.Data, false)

		s.Error(err)
		s.Empty(token)
	})
}
