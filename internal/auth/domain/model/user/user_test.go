package user_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type UserTestSuite struct {
	suite.Suite
}

func TestUserSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestNewUser() {
	type Sut struct {
		Sut  func() (*user.User, error)
		Data *user.User
	}

	makeSut := func() *Sut {
		data := fake.User()

		sut := func() (*user.User, error) {
			return user.NewUser().
				SetEmail(data.Email).
				SetPassword(data.Password).
				SetName(data.Name).
				Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewUser (Error)", func() {
		s.Run("Email", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Email = ""

				user, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(user)
			})
		})

		s.Run("Password", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Password = ""

				user, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(user)
			})
		})

		s.Run("Name", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Name = ""

				user, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(user)
			})
		})
	})

	s.Run("TestNewUser (Success)", func() {
		s.Run("Should return a user when all fields are valid", func() {
			sut := makeSut()

			user, err := sut.Sut()

			s.NoError(err)
			s.NotNil(user)
			s.Equal(sut.Data.Email, user.Email, "should have the same email")
			s.Equal(sut.Data.Password, user.Password, "should have the same password")
			s.Equal(sut.Data.Name, user.Name, "should have the same name")
		})
	})
}
