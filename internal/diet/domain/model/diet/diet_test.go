package diet_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model"
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
			dto := diet.InputDietDTO{
				ID:              data.ID,
				Name:            data.Name,
				Description:     data.Description,
				AllowedFood:     data.AllowedFood,
				RestrictedFood:  data.RestrictedFood,
				DurationInWeeks: data.DurationInWeeks,
				Goal:            data.Goal,
				MealPlan:        data.MealPlan,
				MonthlyCostUSD:  data.MonthlyCostUSD,
			}
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

		s.Run("Should return error when AllowedFood is empty", func() {
			sut := makeSut()
			sut.Data.AllowedFood = nil

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("allowed food cannot be empty", err.Error())
		})

		s.Run("Should return error when RestrictedFood is empty", func() {
			sut := makeSut()
			sut.Data.RestrictedFood = nil

			diet, err := sut.Sut()
			s.Error(err)
			s.Nil(diet)
			s.Equal("restricted food cannot be empty", err.Error())
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
}
