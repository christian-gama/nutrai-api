package fixture

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	persistence "github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	"gorm.io/gorm"
)

type PlanDeps struct {
	Plan *plan.Plan
	Diet *diet.Diet
}

func SavePlan(db *gorm.DB, deps *PlanDeps) *PlanDeps {
	if deps == nil {
		deps = &PlanDeps{}
	}

	if deps.Diet == nil {
		diet := SaveDiet(db, nil)
		deps.Diet = diet.Diet
	}

	plan := deps.Plan
	if plan == nil {
		plan = fake.Plan()
		plan.Diet = deps.Diet

		plan, err := persistence.NewPlan(db).
			Save(context.Background(), repo.SavePlanInput{
				Plan: plan,
			})
		if err != nil {
			panic(fmt.Errorf("could not create plan: %w", err))
		}

		deps.Plan = plan
	}

	return deps
}
