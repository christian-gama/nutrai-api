package manager

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"gorm.io/gorm"
)

// Manager is a generic SQL manager for a model and a schema.
type Manager[Model fmt.Stringer, Schema any] struct {
	*gorm.DB
	model Model
}

// NewManager returns a new Manager.
func NewManager[Model fmt.Stringer, Schema any](db *gorm.DB) *Manager[Model, Schema] {
	var m Model

	return &Manager[Model, Schema]{
		DB:    db,
		model: m,
	}
}

// Save saves the model in the database.
func (m *Manager[Model, Schema]) Save(
	ctx context.Context,
	input SaveInput[Model],
) (*Model, error) {
	db := m.DB.WithContext(ctx)
	var s Schema
	schema := convert.FromModel(&s, &input.Model)

	if err := db.
		Create(&schema).
		Error; err != nil {
		return nil, sql.Error(err, m.model)
	}

	return convert.ToModel(input.Model, &schema), nil
}

// Find finds the model in the database.
func (m *Manager[Model, Schema]) Find(
	ctx context.Context,
	input FindInput[Model],
) (*Model, error) {
	db := m.DB.WithContext(ctx)
	var schema Schema

	if err := db.
		Model(&schema).
		Scopes(
			querying.PreloadScope(input.Preloader),
			querying.FilterScope(input.Filterer),
		).
		Where("id = ?", input.ID).
		First(&schema).
		Error; err != nil {
		return nil, sql.Error(err, m.model)
	}

	var model Model
	return convert.ToModel(&model, &schema), nil
}

// All returns all models in the database.
func (m *Manager[Model, Schema]) All(
	ctx context.Context,
	input AllInput[Model],
) (*queryer.PaginationOutput[*Model], error) {
	db := m.DB.WithContext(ctx)
	var schemas []Schema

	if err := db.
		Scopes(
			querying.PreloadScope(input.Preloader),
			querying.FilterScope(input.Filterer),
			querying.PaginationScope(input.Paginator),
			querying.SortScope(input.Sorter),
		).
		Find(&schemas).
		Error; err != nil {
		return nil, sql.Error(err, m.model)
	}

	var totalCount int64
	var schema Schema
	err := db.
		Model(&schema).
		Scopes(querying.FilterScope(input.Filterer)).
		Count(&totalCount).Error
	if err != nil {
		return nil, sql.Error(err, m.model)
	}

	pagination := &queryer.PaginationOutput[*Model]{}
	for _, sch := range schemas {
		var model Model
		pagination.Results = append(
			pagination.Results,
			convert.ToModel(&model, sch),
		)
	}
	pagination.Total = int(totalCount)

	return pagination, nil
}

// Delete deletes the model in the database.
func (m *Manager[Model, Schema]) Delete(
	ctx context.Context,
	input DeleteInput[Model],
) error {
	db := m.DB.WithContext(ctx)
	var schema Schema

	if err := db.
		Model(&schema).
		Where(input.IDs).
		Delete(&schema).
		Error; err != nil {
		return sql.Error(err, m.model)
	}

	return nil
}

// Update updates the model in the database.
func (m *Manager[Model, Schema]) Update(
	ctx context.Context,
	input UpdateInput[Model],
) error {
	db := m.DB.WithContext(ctx)
	var s Schema
	schema := convert.FromModel(&s, &input.Model)

	if err := db.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Model(&schema).
		Where("id = ?", input.ID).
		Updates(&schema).
		Error; err != nil {
		return sql.Error(err, m.model)
	}

	return nil
}
