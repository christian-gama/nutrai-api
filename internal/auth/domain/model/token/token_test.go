package token_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/token"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type TokenTestSuite struct {
	suite.Suite
}

func TestTokenSuite(t *testing.T) {
	suite.RunUnitTest(t, new(TokenTestSuite))
}

func (s *TokenTestSuite) TestNewToken() {
	type Sut struct {
		Sut  func() (*token.Token, error)
		Data *token.Token
	}

	makeSut := func() *Sut {
		data := fake.Token()

		sut := func() (*token.Token, error) {
			return token.NewToken().
				SetEmail(data.Email).
				SetExpiresAt(data.ExpiresAt).
				SetJti(data.Jti).
				Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewToken (Error)", func() {
		s.Run("Email", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Email = ""

				token, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(token)
			})
		})

		s.Run("Jti", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Jti = ""

				token, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(token)
			})
		})

		s.Run("ExpiresAt", func() {
			s.Run("Should return an error when zero", func() {
				sut := makeSut()
				sut.Data.ExpiresAt = 0

				token, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(token)
			})
		})
	})

	s.Run("TestNewToken (Success)", func() {
		s.Run("Should return a token when all fields are valid", func() {
			sut := makeSut()

			token, err := sut.Sut()

			s.NoError(err)
			s.NotNil(token)
			s.Equal(sut.Data.Email, token.Email, "should have the same email")
			s.Equal(sut.Data.Jti, token.Jti, "should have the same jti")
			s.Equal(sut.Data.ExpiresAt, token.ExpiresAt, "should have the same expiresAt")
		})
	})
}
