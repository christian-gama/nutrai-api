package schema

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
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
	return "restricted_foods"
}
