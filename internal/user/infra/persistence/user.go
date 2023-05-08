package persistence

import (
	"context"

	queryingPort "github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence/schema"
	"github.com/christian-gama/nutrai-api/internal/user/infra/querying"
	"gorm.io/gorm"
)

type userImpl struct {
	db *gorm.DB
}

// NewUser returns a new User.
func NewUser(db *gorm.DB) repo.User {
	return &userImpl{
		db: db,
	}
}

// Save implements repo.User.
func (p *userImpl) Save(
	ctx context.Context,
	input repo.SaveUserInput,
) (value.ID, error) {
	db := p.db.WithContext(ctx)
	userSchema := convert.FromModel(&schema.User{}, &input.User)

	if err := db.
		Create(&userSchema).
		Error; err != nil {
		return 0, sql.Error(err, "user")
	}

	return value.ID(userSchema.ID), nil
}

// Delete implements repo.User.
func (p *userImpl) Delete(ctx context.Context, input repo.DeleteUserInput) error {
	db := p.db.WithContext(ctx)
	var userSchema schema.User

	if err := db.
		Where("id = ?", input.ID).
		Delete(&userSchema).
		Error; err != nil {
		return sql.Error(err, "user")
	}

	return nil
}

// All implements repo.User.
func (p *userImpl) All(
	ctx context.Context,
	input repo.AllUsersInput,
	preload ...string,
) (*queryingPort.PaginationOutput[*user.User], error) {
	db := p.db.WithContext(ctx)
	var users []schema.User

	if err := db.
		Scopes(
			sql.PreloadScope(preload),
			querying.FilterScope(input.Filterer),
			querying.PaginationScope(input.Paginator),
			querying.SortScope(input.Sorter),
		).
		Find(&users).
		Error; err != nil {
		return nil, sql.Error(err, "user")
	}

	var totalCount int64
	err := db.
		Model(&schema.User{}).
		Scopes(querying.FilterScope(input.Filterer)).
		Count(&totalCount).Error
	if err != nil {
		return nil, sql.Error(err, "user")
	}

	pagination := &queryingPort.PaginationOutput[*user.User]{}
	for _, userSchema := range users {
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&user.User{}, userSchema),
		)
	}
	pagination.Total = int(totalCount)

	return pagination, nil
}

// Find implements repo.User.
func (p *userImpl) Find(
	ctx context.Context,
	input repo.FindUserInput,
	preload ...string,
) (*user.User, error) {
	db := p.db.WithContext(ctx)
	var userSchema schema.User

	if err := db.
		Scopes(sql.PreloadScope(preload)).
		Where("id = ?", input.ID).
		First(&userSchema).
		Error; err != nil {
		return nil, sql.Error(err, "user")
	}

	return convert.ToModel(&user.User{}, &userSchema), nil
}

// Update implements repo.User.
func (p *userImpl) Update(
	ctx context.Context,
	input repo.UpdateUserInput,
) error {
	db := p.db.WithContext(ctx)
	userSchema := convert.FromModel(&schema.User{}, &input.User)

	if _, err := p.Find(ctx, repo.FindUserInput{ID: input.User.ID}); err != nil {
		return err
	}

	if err := db.
		Where("id = ?", input.User.ID).
		Updates(&userSchema).
		Error; err != nil {
		return sql.Error(err, "user")
	}

	return nil
}
