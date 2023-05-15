package persistence_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type DietSuite struct {
	suite.SuiteWithConn
	Diet func(db *gorm.DB) repo.Diet
}

func TestDietSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(DietSuite))
}

func (s *DietSuite) SetupTest() {
	s.Diet = func(db *gorm.DB) repo.Diet {
		return persistence.NewDiet(db)
	}
}

func (s *DietSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SaveDietInput) (*diet.Diet, error)
		Ctx   context.Context
		Input repo.SaveDietInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		diet := fake.Diet()
		input := repo.SaveDietInput{
			Diet: diet,
		}

		sut := s.Diet(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should create a new diet", func(db *gorm.DB) {
		sut := makeSut(db)

		diet, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)
		s.NotZero(diet.ID, "Should have an ID")
	})

	s.Run("Should return an error when the diet already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.NoError(err)

		_, err = sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})

	s.Run("Should return an error when the diet is invalid", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.Diet.Description = ""

		_, err := sut.Sut(sut.Ctx, sut.Input)

		s.Error(err)
	})
}
