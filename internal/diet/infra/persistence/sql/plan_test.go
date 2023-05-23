package persistence_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql/schema"
	dietFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	fixture "github.com/christian-gama/nutrai-api/testutils/fixture/diet/sql"

	"github.com/christian-gama/nutrai-api/testutils/suite"
	"gorm.io/gorm"
)

type PlanSuite struct {
	suite.SuiteWithSQLConn
	Plan func(db *gorm.DB) repo.Plan
}

func TestPlanSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(PlanSuite))
}

func (s *PlanSuite) SetupTest() {
	s.Plan = func(db *gorm.DB) repo.Plan {
		return persistence.NewSQLPlan(db)
	}
}

func (s *PlanSuite) TestSave() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.SavePlanInput) (*plan.Plan, error)
		Ctx   context.Context
		Input repo.SavePlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		plan := fake.Plan()
		diet := dietFake.Diet()
		plan.Diet = diet
		input := repo.SavePlanInput{
			Plan: plan,
		}

		sut := s.Plan(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("TestSave (Error)", func(db *gorm.DB) {
		s.Run("Should create a new plan", func(db *gorm.DB) {
			sut := makeSut(db)

			dietDeps := fixture.SaveDiet(db, nil)

			sut.Input.Plan.Diet = dietDeps.Diet

			plan, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(plan.ID, "Should have an ID")
		})
	})

	s.Run("TestSave (Success)", func(db *gorm.DB) {
		s.Run("Should return an error when the plan already exists", func(db *gorm.DB) {
			sut := makeSut(db)

			dietDeps := fixture.SaveDiet(db, nil)

			sut.Input.Plan.Diet = dietDeps.Diet

			plan, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(plan.ID, "Should have an ID")

			plan, err = sut.Sut(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(plan)
		})
	})
}

func (s *PlanSuite) TestFind() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.FindPlanInput) (*plan.Plan, error)
		Ctx   context.Context
		Input repo.FindPlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		plan := fake.Plan()
		diet := dietFake.Diet()
		plan.Diet = diet
		input := repo.FindPlanInput{
			ID: plan.ID,
		}

		sut := s.Plan(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("TestFind (Error)", func(db *gorm.DB) {
		s.Run("Should return an error when the plan does not exist", func(db *gorm.DB) {
			sut := makeSut(db)

			sut.Input.ID = 404_404_404

			plan, err := sut.Sut(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(plan)
		})
	})

	s.Run("TestFind (Success)", func(db *gorm.DB) {
		s.Run("Should return the plan", func(db *gorm.DB) {
			sut := makeSut(db)

			planDeps := fixture.SavePlan(db, nil)

			sut.Input.ID = planDeps.Plan.ID

			plan, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(plan.ID, "Should have an ID")
		})
	})
}

func (s *PlanSuite) TestAll() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.AllPlansInput) (*queryer.PaginationOutput[*plan.Plan], error)
		Ctx   context.Context
		Input repo.AllPlansInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.AllPlansInput{
			Paginator: &querying.Pagination{},
			Sorter:    querying.Sort{},
			Filterer:  querying.Filter{},
		}
		sut := s.Plan(db).All

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("TestAll (Error)", func(db *gorm.DB) {
		s.Run("Should return the correct plans using filters", func(db *gorm.DB) {
			sut := makeSut(db)

			planDeps := fixture.SavePlan(db, nil)
			length := 3

			for i := 0; i < length; i++ {
				fixture.SavePlan(db, nil)
			}

			sut.Input.Filterer = sut.Input.Filterer.Add(
				"ID",
				querying.EqOperator,
				planDeps.Plan.ID,
			)

			result, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(result.Results[0].ID, "Should have an ID")
			s.Equal(1, len(result.Results), "Should have the same length")
			s.Len(result.Results, 1, "Should have the same length")
		})

		s.Run("Should return the correct plans using sorter as desc", func(db *gorm.DB) {
			sut := makeSut(db)

			length := 3

			for i := 0; i < length; i++ {
				fixture.SavePlan(db, nil)
			}

			sut.Input.Sorter = sut.Input.Sorter.Add(
				"ID",
				true,
			)

			result, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Greater(
				int(result.Results[0].ID),
				int(result.Results[1].ID),
				"Should have the correct order",
			)
		})

		s.Run("Should return the correct plans using sorter as asc", func(db *gorm.DB) {
			sut := makeSut(db)

			length := 3

			for i := 0; i < length; i++ {
				fixture.SavePlan(db, nil)
			}

			sut.Input.Sorter = sut.Input.Sorter.Add(
				"ID",
				false,
			)

			result, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Greater(
				int(result.Results[1].ID),
				int(result.Results[0].ID),
				"Should have the correct order",
			)
		})
	})

	s.Run("TestAll (Success)", func(db *gorm.DB) {
		s.Run("Should return all plans", func(db *gorm.DB) {
			sut := makeSut(db)

			length := 3

			for i := 0; i < length; i++ {
				fixture.SavePlan(db, nil)
			}

			result, err := sut.Sut(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotZero(result.Results[0].ID, "Should have an ID")
			s.Equal(length, len(result.Results), "Should have the same length")
			s.Len(result.Results, length, "Should have the same length")
		})
	})
}

func (s *PlanSuite) TestUpdate() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.UpdatePlanInput) error
		Ctx   context.Context
		Input repo.UpdatePlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.UpdatePlanInput{
			Plan: fake.Plan(),
			ID:   1,
		}
		sut := s.Plan(db).Update

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("TestUpdate (Error)", func(db *gorm.DB) {
		s.Run("Should return an error when the plan does not exist", func(db *gorm.DB) {
			sut := makeSut(db)

			sut.Input.ID = 404_404_404

			err := sut.Sut(sut.Ctx, sut.Input)

			s.Error(err)
		})
	})

	s.Run("TestUpdate (Success)", func(db *gorm.DB) {
		s.Run("Should update the plan", func(db *gorm.DB) {
			sut := makeSut(db)

			planDeps := fixture.SavePlan(db, nil)

			*sut.Input.Plan = *planDeps.Plan
			sut.Input.Plan.Text = "new text"
			sut.Input.ID = planDeps.Plan.ID
			sut.Input.Plan.ID = planDeps.Plan.ID

			err := sut.Sut(sut.Ctx, sut.Input)

			s.Require().NoError(err)
			s.HasChanged(planDeps.Plan, sut.Input.Plan)
		})
	})
}

func (s *PlanSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeletePlanInput) error
		Ctx   context.Context
		Input repo.DeletePlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := repo.DeletePlanInput{}
		sut := s.Plan(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("TestDelete (Success)", func(db *gorm.DB) {
		s.Run("Should delete the plan", func(db *gorm.DB) {
			sut := makeSut(db)

			planDeps := fixture.SavePlan(db, nil)

			sut.Input.IDs = []value.ID{planDeps.Plan.ID}

			err := sut.Sut(sut.Ctx, sut.Input)

			s.Require().NoError(err)
			s.SQLRecordDoesNotExist(db, &schema.Plan{})
		})
	})
}
