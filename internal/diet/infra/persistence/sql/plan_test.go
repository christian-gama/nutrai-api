package persistence_test

import (
	"context"
	"sort"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql/schema"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	dietFixture "github.com/christian-gama/nutrai-api/testutils/fixture/diet/sql"
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
		Input *repo.SavePlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		plan := fake.Plan()

		input := &repo.SavePlanInput{
			Plan: plan,
		}

		sut := s.Plan(db).Save

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
		}
	}

	s.Run("Should create a new plan", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := dietFixture.SaveDiet(db, nil)
		sut.Input.Plan.DietID = dietDeps.Diet.ID
		plan, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotZero(plan.ID, "Should have an ID")
	})

	s.Run("Should return an error when the plan already exists", func(db *gorm.DB) {
		sut := makeSut(db)

		dietDeps := dietFixture.SaveDiet(db, nil)
		sut.Input.Plan.DietID = dietDeps.Diet.ID
		_, err := sut.Sut(sut.Ctx, *sut.Input)
		s.NoError(err)
		s.SQLRecordExist(db, &schema.Plan{})

		_, err = sut.Sut(sut.Ctx, *sut.Input)
		s.Error(err)
	})

	s.Run("Should return an error when the diet does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.Plan.ID = 404_404_404
		_, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *PlanSuite) TestDelete() {
	type Sut struct {
		Sut   func(ctx context.Context, input repo.DeletePlanInput) error
		Ctx   context.Context
		Input *repo.DeletePlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := &repo.DeletePlanInput{}
		sut := s.Plan(db).Delete

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should delete a plan", func(db *gorm.DB) {
		sut := makeSut(db)

		planDeps := fixture.SavePlan(db, nil)

		sut.Input.IDs = []value.ID{planDeps.Plan.ID}

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.SQLRecordDoesNotExist(db, &schema.Plan{})
	})

	s.Run("Should delete nothing if the plan ID is invalid", func(db *gorm.DB) {
		sut := makeSut(db)

		planDeps := fixture.SavePlan(db, nil)

		sut.Input.IDs = []value.ID{planDeps.Plan.ID + 1}

		err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.SQLRecordExist(db, &schema.Plan{})
	})
}

func (s *PlanSuite) TestFind() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.FindPlanInput,
		) (*plan.Plan, error)
		Ctx   context.Context
		Input *repo.FindPlanInput
	}

	makeSut := func(db *gorm.DB) Sut {
		ctx := context.Background()
		input := &repo.FindPlanInput{
			ID: 0,
		}
		sut := s.Plan(db).Find

		return Sut{
			Sut:   sut,
			Ctx:   ctx,
			Input: input,
		}
	}

	s.Run("Should find a plan", func(db *gorm.DB) {
		sut := makeSut(db)

		planDeps := fixture.SavePlan(db, nil)

		sut.Input.ID = planDeps.Plan.ID

		plan, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(plan.ID, planDeps.Plan.ID, "Should have the same ID")
	})

	s.Run("Should return an error if the plan does not exist", func(db *gorm.DB) {
		sut := makeSut(db)

		sut.Input.ID = 404_404_404

		_, err := sut.Sut(sut.Ctx, *sut.Input)

		s.Error(err)
	})
}

func (s *PlanSuite) TestAll() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.AllPlansInput,
		) (*queryer.PaginationOutput[*plan.Plan], error)
		Ctx   context.Context
		Input *repo.AllPlansInput
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
			Input: &input,
		}
	}

	s.Run("Should find all plans", func(db *gorm.DB) {
		sut := makeSut(db)

		length := 3
		for i := 0; i < length; i++ {
			fixture.SavePlan(db, nil)
		}

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.NotZero(result.Results[0].ID, "Should have a valid id")
		s.Equal(length, result.Total, "Should return %d total", length)
		s.Len(result.Results, length, "Should return %d results", length)
	})

	s.Run("Should return the correct plans using filter", func(db *gorm.DB) {
		sut := makeSut(db)

		plan := fake.Plan()
		plan.Text = "test"
		planDeps := fixture.SavePlan(db, &fixture.PlanDeps{Plan: plan})
		length := 2
		for i := 0; i < length; i++ {
			plan := fake.Plan()
			fixture.SavePlan(db, &fixture.PlanDeps{Plan: plan})
		}

		sut.Input.Filterer = sut.Input.Filterer.Add(
			"text",
			querying.EqOperator,
			plan.Text,
		)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(result.Results[0].ID, planDeps.Plan.ID, "Should have the same ID")
		s.Equal(1, result.Total, "Should return only one plan")
		s.Len(result.Results, 1, "Should return only one plan")
	})

	s.Run("Should return the correct plans using sorter as desc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SavePlan(db, nil)
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

	s.Run("Should return the correct plans using sorter as asc", func(db *gorm.DB) {
		sut := makeSut(db)

		for i := 0; i < 3; i++ {
			fixture.SavePlan(db, nil)
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

	s.Run("Should return the correct plans using pagination", func(db *gorm.DB) {
		sut := makeSut(db)

		plans := make([]*plan.Plan, 0)
		for i := 0; i < 3; i++ {
			planDeps := fixture.SavePlan(db, nil)
			plans = append(plans, planDeps.Plan)
		}

		sort.Slice(plans, func(i, j int) bool {
			return plans[i].ID >= plans[j].ID
		})

		sut.Input.Paginator = sut.Input.Paginator.SetLimit(1).SetPage(1)
		sut.Input.Sorter = sut.Input.Sorter.Add("id", true)

		result, err := sut.Sut(sut.Ctx, *sut.Input)

		s.NoError(err)
		s.Equal(3, result.Total, "Should return the correct total")
		s.Len(result.Results, 1, "Should return the correct number of plans")
		s.Equal(int(plans[0].ID), int(result.Results[0].ID), "Should return the correct plan")
	})
}

func (s *PlanSuite) TestUpdate() {
	type Sut struct {
		Sut func(
			ctx context.Context,
			input repo.UpdatePlanInput,
		) error
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

	s.Run("Should update a plan", func(db *gorm.DB) {
		sut := makeSut(db)

		planDeps := fixture.SavePlan(db, nil)

		*sut.Input.Plan = *planDeps.Plan
		sut.Input.Plan.Text = "new text"
		sut.Input.ID = planDeps.Plan.ID

		err := sut.Sut(sut.Ctx, sut.Input)

		s.Require().NoError(err)
		s.HasChanged(planDeps.Plan, sut.Input.Plan)
	})
}
