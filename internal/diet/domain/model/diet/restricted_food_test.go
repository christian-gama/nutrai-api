package diet_test

import (
	"testing"

	diet "github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
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
		Sut  func() (*diet.RestrictedFood, error)
		Data *diet.RestrictedFood
	}

	makeSut := func() *Sut {
		data := fake.RestrictedFood()

		sut := func() (*diet.RestrictedFood, error) {
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewRestrictedFood (Error)", func() {
		s.Run("Name", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Name = ""

				restrictedfood, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(restrictedfood)
			})
		})
	})
}
