package persistence_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/user/sql"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type DietSuite struct {
	suite.SuiteWithSQLConn
	Diet func(db *gorm.DB) repo.Diet
}

func TestDietSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(DietSuite))
}

func (s *DietSuite) SetupTest() {
	s.Diet = func(db *gorm.DB) repo.Diet {
		return persistence.NewSQLDiet(db)
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

	s.Run("TestSave (Error)", func(db *gorm.DB) {
		s.Run("Should return an error when the diet already exists", func(db *gorm.DB) {
			sut := makeSut(db)

			patientDeps := fixture.SavePatient(db, nil)

			sut.Input.Diet.PatientID = patientDeps.Patient.ID

			_, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)

			_, err = sut.Sut(sut.Ctx, sut.Input)

			s.Error(err)
		})
	})

	s.Run("TestSave (Success)", func(db *gorm.DB) {
		s.Run("Should create a new diet", func(db *gorm.DB) {
			sut := makeSut(db)

			patientDeps := fixture.SavePatient(db, nil)

			sut.Input.Diet.PatientID = patientDeps.Patient.ID

			diet, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(diet.ID, "Should have an ID")
		})
	})
}
