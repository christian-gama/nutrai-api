package persistence

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql/manager"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence/sql/schema"
	"gorm.io/gorm"
)

// userImpl is the implementation of repo.User.
type userImpl struct {
	manager *manager.Manager[user.User, schema.User]
}

// NewUser returns a new User.
func NewUser(db *gorm.DB) repo.User {
	return &userImpl{
		manager: manager.NewManager[user.User, schema.User](db),
	}
}

// All implements repo.User.
func (p *userImpl) All(ctx context.Context, input repo.AllUsersInput) (*querying.PaginationOutput[*user.User], error) {
	return p.manager.All(ctx,
		manager.AllInput[user.User]{
			Filterer:  input.Filterer,
			Paginator: input.Paginator,
			Sorter:    input.Sorter,
			Preloader: input.Preloader,
		},
	)
}

// Delete implements repo.User.
func (p *userImpl) Delete(ctx context.Context, input repo.DeleteUserInput) error {
	return p.manager.Delete(ctx,
		manager.DeleteInput[user.User]{
			IDs: input.IDs,
		},
	)
}

// Find implements repo.User.
func (p *userImpl) Find(ctx context.Context, input repo.FindUserInput) (*user.User, error) {
	return p.manager.Find(ctx,
		manager.FindInput[user.User]{
			ID:        input.ID,
			Filterer:  input.Filterer,
			Preloader: input.Preloader,
		},
	)
}

// FindByEmail implements repo.User.
func (p *userImpl) FindByEmail(ctx context.Context, input repo.FindByEmailUserInput) (*user.User, error) {
	db := p.manager.WithContext(ctx)
	var schema schema.User

	if err := db.
		Model(&schema).
		Where("email = ?", input.Email).
		First(&schema).
		Error; err != nil {
		return nil, sql.Error(err, schema.TableName())
	}

	return convert.ToModel(&user.User{}, &schema), nil
}

// Save implements repo.User.
func (p *userImpl) Save(ctx context.Context, input repo.SaveUserInput) (*user.User, error) {
	return p.manager.Save(ctx,
		manager.SaveInput[user.User]{
			Model: input.User,
		},
	)
}

// Update implements repo.User.
func (p *userImpl) Update(ctx context.Context, input repo.UpdateUserInput) error {
	return p.manager.Update(ctx,
		manager.UpdateInput[user.User]{
			Model: input.User,
			ID:    input.ID,
		},
	)
}
