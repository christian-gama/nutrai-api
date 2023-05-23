package persistence

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/sql/schema"
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/manager"
	"gorm.io/gorm"
)

// userSQLImpl is the SQL implementation of repo.User.
type userSQLImpl struct {
	manager *manager.Manager[user.User, schema.User]
}

// NewSQLUser returns a new User.
func NewSQLUser(db *gorm.DB) repo.User {
	if db == nil {
		panic(errors.New("db cannot be nil"))
	}

	return &userSQLImpl{
		manager: manager.NewManager[user.User, schema.User](db),
	}
}

// All implements repo.User.
func (p *userSQLImpl) All(
	ctx context.Context,
	input repo.AllUsersInput,
) (*queryer.PaginationOutput[*user.User], error) {
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
func (p *userSQLImpl) Delete(ctx context.Context, input repo.DeleteUserInput) error {
	return p.manager.Delete(ctx,
		manager.DeleteInput[user.User]{
			IDs: input.IDs,
		},
	)
}

// Find implements repo.User.
func (p *userSQLImpl) Find(ctx context.Context, input repo.FindUserInput) (*user.User, error) {
	return p.manager.Find(ctx,
		manager.FindInput[user.User]{
			ID:        input.ID,
			Filterer:  input.Filterer,
			Preloader: input.Preloader,
		},
	)
}

// FindByEmail implements repo.User.
func (p *userSQLImpl) FindByEmail(
	ctx context.Context,
	input repo.FindByEmailUserInput,
) (*user.User, error) {
	db := p.manager.WithContext(ctx)
	var schema schema.User
	var model user.User

	if err := db.
		Model(&schema).
		Where("email = ?", input.Email).
		First(&schema).
		Error; err != nil {
		return nil, sql.Error(err, model)
	}

	return convert.ToModel(&model, &schema), nil
}

// Save implements repo.User.
func (p *userSQLImpl) Save(ctx context.Context, input repo.SaveUserInput) (*user.User, error) {
	return p.manager.Save(ctx,
		manager.SaveInput[user.User]{
			Model: input.User,
		},
	)
}

// Update implements repo.User.
func (p *userSQLImpl) Update(ctx context.Context, input repo.UpdateUserInput) error {
	return p.manager.Update(ctx,
		manager.UpdateInput[user.User]{
			Model: input.User,
			ID:    input.ID,
		},
	)
}
