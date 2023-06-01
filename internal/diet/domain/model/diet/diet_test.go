package diet_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type DietTestSuite struct {
	suite.Suite
}

func TestDietSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DietTestSuite))
}

func (s *DietTestSuite) TestNewDiet() {
	type Sut struct {
		Sut  func() (*diet.Diet, error)
		Data *diet.Diet
	}

	makeSut := func() *Sut {
		data := fake.Diet()

		sut := func() (*diet.Diet, error) {
			return data.Validate()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewDiet (Error)", func() {

		s.Run("MonthlyCostUSD", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.MonthlyCostUSD = 0

				diet, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(diet)
			})

			s.Run("Should return an error when negative", func() {
				sut := makeSut()
				sut.Data.MonthlyCostUSD = -1

				diet, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(diet)
			})
		})

		s.Run("Description", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Description = ""

				diet, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(diet)
			})
		})

		s.Run("DurationInWeeks", func() {
			s.Run("Should return an error when less than min", func() {
				sut := makeSut()
				sut.Data.DurationInWeeks = 0

				diet, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(diet)
			})
		})

		s.Run("Goal", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Goal = ""

				diet, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(diet)
			})

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.Goal = "invalid"

				diet, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(diet)
			})
		})

		s.Run("MealPlan", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.MealPlan = ""

				diet, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(diet)
			})

			s.Run("Should return an error when invalid", func() {
				sut := makeSut()
				sut.Data.MealPlan = "invalid"

				diet, err := sut.Sut()

				s.ErrorAsInvalid(err)
				s.Nil(diet)
			})
		})

		s.Run("Name", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Name = ""

				diet, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(diet)
			})
		})
	})

	s.Run("TestNewDiet (Success)", func() {
		s.Run("Should return a diet when all fields are valid", func() {
			sut := makeSut()

			diet, err := sut.Sut()

			s.NoError(err)
			s.NotNil(diet)
			s.Equal(sut.Data.Name, diet.Name, "should have the same name")
			s.Equal(sut.Data.Description, diet.Description, "should have the same description")
			// s.Equal(
			// 	sut.Data.RestrictedFood,
			// 	diet.RestrictedFood,
			// 	"should have the same restricted food",
			// )
			s.Equal(
				sut.Data.DurationInWeeks,
				diet.DurationInWeeks,
				"should have the same duration in weeks",
			)
			s.Equal(sut.Data.Goal, diet.Goal, "should have the same goal")
			s.Equal(sut.Data.MealPlan, diet.MealPlan, "should have the same meal plan")
			s.Equal(
				sut.Data.MonthlyCostUSD,
				diet.MonthlyCostUSD,
				"should have the same monthly cost in USD",
			)
		})
	})
}
