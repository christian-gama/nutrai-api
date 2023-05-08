package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/manager"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence/schema"
	"gorm.io/gorm"
)

type userImpl struct {
	manager *manager.Manager[user.User, schema.User]
}

// All implements repo.User.
func (p *userImpl) All(ctx context.Context, input repo.AllUsersInput, preload ...string) (*querying.PaginationOutput[*user.User], error) {
	return p.manager.All(ctx, manager.AllInput[user.User]{Filterer: input.Filterer, Paginator: input.Paginator, Sorter: input.Sorter}, preload...)
}

// Delete implements repo.User.
func (p *userImpl) Delete(ctx context.Context, input repo.DeleteUserInput) error {
	return p.manager.Delete(ctx, manager.DeleteInput[user.User]{IDs: input.IDs})
}

// Find implements repo.User.
func (p *userImpl) Find(ctx context.Context, input repo.FindUserInput, preload ...string) (*user.User, error) {
	return p.manager.Find(ctx, manager.FindInput[user.User]{ID: input.ID, Filterer: input.Filterer}, preload...)
}

// Save implements repo.User.
func (p *userImpl) Save(ctx context.Context, input repo.SaveUserInput) (*user.User, error) {
	return p.manager.Save(ctx, manager.SaveInput[user.User]{Model: input.User})
}

// Update implements repo.User.
func (p *userImpl) Update(ctx context.Context, input repo.UpdateUserInput) error {
	return p.manager.Update(ctx, manager.UpdateInput[user.User]{Model: input.User, ID: input.ID})
}

// NewUser returns a new User.
func NewUser(db *gorm.DB) repo.User {
	return &userImpl{
		manager: manager.NewManager[user.User, schema.User](db),
	}
}
