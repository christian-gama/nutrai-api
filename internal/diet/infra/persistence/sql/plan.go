package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/sql/schema"
	"gorm.io/gorm"
)

// planImpl is the implementation of repo.Plan.
type planSQLImpl struct {
	manager *manager.Manager[plan.Plan, schema.Plan]
}

// All implements repo.Plan.
func (p *planSQLImpl) All(ctx context.Context, input repo.AllPlansInput) (*queryer.PaginationOutput[*plan.Plan], error) {
	return p.manager.All(ctx, manager.AllInput[plan.Plan]{Filterer: input.Filterer, Paginator: input.Paginator, Sorter: input.Sorter, Preloader: input.Preloader})
}

// Delete implements repo.Plan.
func (p *planSQLImpl) Delete(ctx context.Context, input repo.DeletePlanInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[plan.Plan]{IDs: input.IDs})
}

// Find implements repo.Plan.
func (p *planSQLImpl) Find(ctx context.Context, input repo.FindPlanInput) (*plan.Plan, error) {
	return p.manager.Find(ctx, manager.FindInput[plan.Plan]{ID: input.ID, Filterer: input.Filterer, Preloader: input.Preloader})
}

// Save implements repo.Plan.
func (p *planSQLImpl) Save(ctx context.Context, input repo.SavePlanInput) (*plan.Plan, error) {
	return p.manager.Save(ctx, manager.SaveInput[plan.Plan]{Model: input.Plan})
}

// Update implements repo.Plan.
func (p *planSQLImpl) Update(ctx context.Context, input repo.UpdatePlanInput) error {
	return p.manager.Update(ctx, manager.UpdateInput[plan.Plan]{Model: input.Plan})
}

// NewSQLPlan returns a new instance of repo.Plan.
func NewSQLPlan(db *gorm.DB) repo.Plan {
	return &planSQLImpl{
		manager: manager.NewManager[plan.Plan, schema.Plan](db),
	}
}
