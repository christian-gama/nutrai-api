package restrictedfood_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/restrictedfood"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/restrictedfood"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type RestrictedFoodTestSuite struct {
	suite.Suite
}

func TestRestrictedFoodSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RestrictedFoodTestSuite))
}

func (s *RestrictedFoodTestSuite) TestNewRestrictedFood() {
	type Sut struct {
		Sut  func() (*restrictedfood.RestrictedFood, error)
		Data *restrictedfood.RestrictedFood
	}

	makeSut := func() *Sut {
		data := fake.RestrictedFood()

		sut := func() (*restrictedfood.RestrictedFood, error) {
			return restrictedfood.NewBuilder().
				SetID(data.ID).
				SetName(data.Name).
				Build()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewRestrictedFood (Error)", func() {
		s.Run("ID", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.ID = 0

				restrictedfood, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(restrictedfood)
			})
		})

		s.Run("Name", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Name = ""

				restrictedfood, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(restrictedfood)
			})

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.Name = "invalid"

				restrictedfood, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(restrictedfood)
			})
		})
	})
}
