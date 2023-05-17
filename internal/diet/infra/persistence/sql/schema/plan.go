package schema

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
)

// Plan is the database schema for plans.
type Plan struct {
	ID     coreValue.ID `gorm:"primaryKey"`
	DietID coreValue.ID
	Diet   *Diet `gorm:"foreignKey:DietID"`
	Text   value.Plan
}

// TableName returns the name of the table.
func (Plan) TableName() string {
	return "plans"
}
