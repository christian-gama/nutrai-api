package schema

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/table"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/module"
)

// RestrictedFood is the database schema for restricted foods.
type RestrictedFood struct {
	ID     coreValue.ID `gorm:"primaryKey"`
	DietID coreValue.ID
	Diet   *Diet `gorm:"foreignKey:DietID"`
	Name   value.Name
}

// TableName returns the table name for the RestrictedFood schema.
func (RestrictedFood) TableName() string {
	return table.Name(module.Module, "restricted_foods")
}
