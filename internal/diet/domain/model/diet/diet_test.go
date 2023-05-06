package diet_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
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
			dto := diet.DietInput(*data)
			return diet.NewDiet(dto)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewDiet (Error)", func() {
		s.Run("Should return error when ID is empty", func() {
			sut := makeSut()
			sut.Data.ID = 0

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid ID", err.Error())
		})

		s.Run("Should return error when Name is empty", func() {
			sut := makeSut()
			sut.Data.Name = ""

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid name", err.Error())
		})

		s.Run("Should return error when Description is empty", func() {
			sut := makeSut()
			sut.Data.Description = ""

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid description", err.Error())
		})

		s.Run("Should return error when a restricted food is not valid", func() {
			sut := makeSut()
			stringWith100Characters := ""

			for i := 0; i < 100; i++ {
				stringWith100Characters += "a"
			}

			sut.Data.RestrictedFood = []value.RestrictedFood{value.RestrictedFood(stringWith100Characters)}

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid restricted food", err.Error())
		})

		s.Run("Should return error when DurationInWeeks is empty", func() {
			sut := makeSut()
			sut.Data.DurationInWeeks = 0

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid duration in weeks", err.Error())
		})

		s.Run("Should return error when Goal is empty", func() {
			sut := makeSut()
			sut.Data.Goal = ""

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid goal", err.Error())
		})

		s.Run("Should return error when MealPlan is empty", func() {
			sut := makeSut()
			sut.Data.MealPlan = ""

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid meal plan", err.Error())
		})

		s.Run("Should return error when MonthlyCostUSD is empty", func() {
			sut := makeSut()
			sut.Data.MonthlyCostUSD = 0

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("invalid monthly cost", err.Error())
		})
	})

	s.Run("TestNewDiet (Success)", func() {
		sut := makeSut()
		diet, err := sut.Sut()

		s.NoError(err)
		s.NotNil(diet)
		s.Equal(sut.Data.ID, diet.ID)
		s.Equal(sut.Data.Name, diet.Name)
		s.Equal(sut.Data.Description, diet.Description)
		s.Equal(sut.Data.RestrictedFood, diet.RestrictedFood)
		s.Equal(sut.Data.DurationInWeeks, diet.DurationInWeeks)
		s.Equal(sut.Data.Goal, diet.Goal)
		s.Equal(sut.Data.MealPlan, diet.MealPlan)
		s.Equal(sut.Data.MonthlyCostUSD, diet.MonthlyCostUSD)
	})
}
