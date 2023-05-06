package user_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
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
			dto := user.UserInput(*data)
			return user.New(&dto)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewUser (Error)", func() {
		s.Run("ID", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.ID = 0

				user, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(user)
			})
		})

		s.Run("email", func() {
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

		s.Run("password", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Password = ""

				user, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(user)
			})

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.Password = "123"

				user, err := sut.Sut()

				s.ErrorAsInvalid(err)
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
			s.Equal(sut.Data.ID, user.ID)
			s.Equal(sut.Data.Email, user.Email)
			s.Equal(sut.Data.Password, user.Password)
		})
	})
}
