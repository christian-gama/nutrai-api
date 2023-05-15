package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/diet/infra/persistence/schema"
	"gorm.io/gorm"
)

// dietImpl is the implementation of repo.Diet.
type dietImpl struct {
	manager *manager.Manager[diet.Diet, schema.Diet]
}

// All implements repo.Diet
func (p *dietImpl) All(ctx context.Context, input repo.AllDietsInput) (*queryer.PaginationOutput[*diet.Diet], error) {
	return p.manager.All(ctx, manager.AllInput[diet.Diet]{Filterer: input.Filterer, Paginator: input.Paginator, Sorter: input.Sorter, Preloader: input.Preloader})
}

// Delete implements repo.Diet
func (p *dietImpl) Delete(ctx context.Context, input repo.DeleteDietInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[diet.Diet]{IDs: input.IDs})
}

// Find implements repo.Diet
func (p *dietImpl) Find(ctx context.Context, input repo.FindDietInput) (*diet.Diet, error) {
	return p.manager.Find(ctx, manager.FindInput[diet.Diet]{ID: input.ID, Filterer: input.Filterer, Preloader: input.Preloader})
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
