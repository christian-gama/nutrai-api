package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/infra/persistence/schema"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/manager"
	"gorm.io/gorm"
)

// dietImpl is the implementation of repo.Diet.
type dietImpl struct {
	manager *manager.Manager[diet.Diet, schema.Diet]
}

// All implements repo.Diet
func (p *dietImpl) All(ctx context.Context, input repo.AllDietsInput, preload ...string) (*querying.PaginationOutput[*diet.Diet], error) {
	return p.manager.All(ctx, manager.AllInput[diet.Diet]{Filterer: input.Filterer, Paginator: input.Paginator, Sorter: input.Sorter}, preload...)
}

// Delete implements repo.Diet
func (p *dietImpl) Delete(ctx context.Context, input repo.DeleteDietInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[diet.Diet]{IDs: input.IDs})
}

// Find implements repo.Diet
func (p *dietImpl) Find(ctx context.Context, input repo.FindDietInput, preload ...string) (*diet.Diet, error) {
	return p.manager.Find(ctx, manager.FindInput[diet.Diet]{ID: input.ID, Filterer: input.Filterer}, preload...)
}

// Save implements repo.Diet
func (p *dietImpl) Save(ctx context.Context, input repo.SaveDietInput) (*diet.Diet, error) {
	return p.manager.Save(ctx, manager.SaveInput[diet.Diet]{Model: input.Diet})
}

// Update implements repo.Diet
func (p *dietImpl) Update(ctx context.Context, input repo.UpdateDietInput) error {
	return p.manager.Update(ctx, manager.UpdateInput[diet.Diet]{Model: input.Diet})
}

// NewDiet returns a new instance of repo.Diet.
func NewDiet(db *gorm.DB) repo.Diet {
	return &dietImpl{
		manager: manager.NewManager[diet.Diet, schema.Diet](db),
	}
}
