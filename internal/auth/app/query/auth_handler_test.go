package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/query"
	jwtFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/jwt"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	jwtMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/jwt"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/auth/domain/repo"
)

type AuthHandlerSuite struct {
	suite.Suite
}

func TestAuthHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AuthHandlerSuite))
}

func (s *AuthHandlerSuite) TestAuthHandler() {
	type Mock struct {
		UserRepo *userRepoMock.User
		Verifier *jwtMock.Verifier
	}

	type Sut struct {
		Sut   query.AuthHandler
		Ctx   context.Context
		Input *query.AuthInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			UserRepo: userRepoMock.NewUser(s.T()),
			Verifier: jwtMock.NewVerifier(s.T()),
		}

		input := fake.AuthInput()

		sut := query.NewAuthHandler(mock.UserRepo, mock.Verifier)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return a AuthOutput", func() {
		sut := makeSut()

		user := userFake.User()
		claims := jwtFake.AccessTokenClaims()
		claims.Sub.Email = user.Email

		sut.Mock.Verifier.
			On("Verify", sut.Input.Access, false).
			Return(claims, nil)

		sut.Mock.UserRepo.
			On("FindByEmail", sut.Ctx, mock.Anything).
			Return(user, nil)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Equal(user.Email, output.Email, "Email should be the same")
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Access, false).
			Return(jwtFake.AccessTokenClaims(), nil)

		sut.Mock.UserRepo.
			On("FindByEmail", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorAsUnauthorized(err)
	})

	s.Run("Should return an error when the verifier fails", func() {
		sut := makeSut()

		sut.Mock.Verifier.
			On("Verify", sut.Input.Access, false).
			Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorAsUnauthorized(err)
	})
}
