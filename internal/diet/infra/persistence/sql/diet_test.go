package persistence_test

import (
	"context"
	"sort"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql/schema"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/diet/sql"
	patientFixture "github.com/christian-gama/nutrai-api/testutils/fixture/patient/sql"
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
		Input *repo.SaveDietInput
	}

	makeSut := func(db *gorm.DB) Sut {
		diet := fake.Diet()

		input := &repo.SaveDietInput{
			Diet: diet,
		}

		sut := s.Diet(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
		}
	}

	s.Run("Should create a new diet", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := patientFixture.SavePatient(db, nil)
		sut.Input.Diet.PatientID = patientDeps.Patient.ID
		diet, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotZero(diet.ID, "Should have an ID")
	})

	s.Run("Should return an error when the diet already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		patientDeps := patientFixture.SavePatient(db, nil)
		sut.Input.Diet.PatientID = patientDeps.Patient.ID
		_, err := sut.Sut(sut.Ctx, *sut.Input)
		s.NoError(err)
		s.SQLRecordExist(db, &schema.Diet{})

		_, err = sut.Sut(sut.Ctx, *sut.Input)
		s.Error(err)
	})

	s.Run("Should return an error when the patient does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.Diet.ID = 404_404_404
		_, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *DietSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeleteDietInput) error
		Ctx   context.Context
		Input *repo.DeleteDietInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := &repo.DeleteDietInput{}
		sut := s.Diet(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete a diet", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := fixture.SaveDiet(db, nil)

		sut.Input.IDs = []value.ID{dietDeps.Diet.ID}

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.SQLRecordDoesNotExist(db, &schema.Diet{})
	})

	s.Run("Should delete nothing if the diet ID is invalid", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := fixture.SaveDiet(db, nil)

		sut.Input.IDs = []value.ID{dietDeps.Diet.ID + 1}

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.SQLRecordExist(db, &schema.Diet{})
	})
}

func (s *DietSuite) TestFind() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.FindDietInput,
		) (*diet.Diet, error)
		Ctx   context.Context
		Input *repo.FindDietInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := &repo.FindDietInput{
			ID: 0,
		}
		sut := s.Diet(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find a diet", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := fixture.SaveDiet(db, nil)

		sut.Input.ID = dietDeps.Diet.ID

		diet, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(diet.ID, dietDeps.Diet.ID, "Should have the same ID")
	})

	s.Run("Should return an error if the diet does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *DietSuite) TestAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.AllDietsInput,
		) (*queryer.PaginationOutput[*diet.Diet], error)
		Ctx   context.Context
		Input *repo.AllDietsInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.AllDietsInput{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Diet(db).All

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: &input,
		}
	}

	s.Run("Should find all diets", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.SaveDiet(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "Should have a valid id")
		s.Equal(length, result.Total, "Should return %d total", length)
		s.Len(result.Results, length, "Should return %d results", length)
	})

	s.Run("Should return the correct diets using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		diet := fake.Diet()
		diet.Name = "test"
		dietDeps := fixture.SaveDiet(db, &fixture.DietDeps{Diet: diet})
		length := 2
		for i := 0; i < length; i++ {
			diet := fake.Diet()
			fixture.SaveDiet(db, &fixture.DietDeps{Diet: diet})
		}

		sut.Input.Filterer = sut.Input.Filterer.Add(
			"name",
			querying.EqOperator,
			diet.Name,
		)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(result.Results[0].ID, dietDeps.Diet.ID, "Should have the same ID")
		s.Equal(1, result.Total, "Should return only one diet")
		s.Len(result.Results, 1, "Should return only one diet")
	})

	s.Run("Should return the correct diets using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SaveDiet(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Greater(
			int(result.Results[1].ID),
			int(result.Results[2].ID),
			"Should have the correct order",
		)
	})

	s.Run("Should return the correct diets using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SaveDiet(db, nil)
		}

		sut.Input.Sorter = sut.Input.Sorter.Add("id", false)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Greater(
			int(result.Results[2].ID),
			int(result.Results[1].ID),
			"Should have the correct order",
		)
	})

	s.Run("Should return the correct diets using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		diets := make([]*diet.Diet, 0)
		for i := 0; i < 3; i++ {
			dietDeps := fixture.SaveDiet(db, nil)
			diets = append(diets, dietDeps.Diet)
		}

		sort.Slice(diets, func(i, j int) bool {
			return diets[i].ID >= diets[j].ID
		})

		sut.Input.Paginator = sut.Input.Paginator.SetLimit(1).SetPage(1)
		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(3, result.Total, "Should return the correct total")
		s.Len(result.Results, 1, "Should return the correct number of diets")
		s.Equal(int(diets[0].ID), int(result.Results[0].ID), "Should return the correct diet")
	})

	s.Run("Should return the correct diets using preloads", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := fixture.SaveDiet(db, nil)
		fixture.SavePlan(db, &fixture.PlanDeps{
			Diet: dietDeps.Diet,
		})

		sut.Input.Preloader = querying.AddPreload("plans").Add("patient")

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(
			dietDeps.Diet.ID,
			result.Results[0].Plans[0].DietID,
			"Should return the correct plan",
		)
		s.Len(result.Results[0].Plans, 1, "Should return the correct number of plans")
	})
}

func (s *DietSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.UpdateDietInput,
		) error
		Ctx   context.Context
		Input repo.UpdateDietInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.UpdateDietInput{
			Diet: fake.Diet(),
			ID:   1,
		}
		sut := s.Diet(db).Update

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should update a diet", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := fixture.SaveDiet(db, nil)

		*sut.Input.Diet = *dietDeps.Diet
		sut.Input.Diet.Name = "new name"
		sut.Input.ID = dietDeps.Diet.ID

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Require().NoError(err)
		s.HasChanged(dietDeps.Diet, sut.Input.Diet)
	})
}
