package user_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/user"
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
			return user.NewBuilder().
				SetEmail(data.Email).
				SetPassword(data.Password).
				SetName(data.Name).
				Build()
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

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.Email = "invalid"

				user, err := sut.Sut()

				s.ErrorAsInvalid(err)
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

			s.Run("Should return an error when too short", func() {
				sut := makeSut()
				sut.Data.Name = "a"

				user, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(user)
			})

			s.Run("Should return an error when too long", func() {
				sut := makeSut()
				sut.Data.Name = value.Name(strings.Repeat("a", 256))

				user, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(user)
			})
		})

		s.Run("Should return multiple errors when multiple fields are invalid", func() {
			sut := makeSut()
			sut.Data.Email = ""
			sut.Data.Password = ""
			sut.Data.Name = ""

			user, err := sut.Sut()

			e := err.(*errutil.Error)
			s.Equal(3, e.Len(), "should have 3 errors")
			s.Nil(user)
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
