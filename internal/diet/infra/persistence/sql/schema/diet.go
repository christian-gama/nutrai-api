package schema

import (
	"time"

	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/table"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/module"
	"github.com/christian-gama/nutrai-api/internal/patient/infra/persistence/sql/schema"
)

// Diet is the database schema for diets.
type Diet struct {
	ID              coreValue.ID `gorm:"primaryKey"`
	PatientID       coreValue.ID
	Patient         *schema.Patient `gorm:"foreignKey:PatientID"`
	Name            value.Name
	Description     value.Description
	DurationInWeeks value.DurationInWeeks
	Goal            value.Goal
	MealPlan        value.MealPlan
	MonthlyCostUSD  value.MonthlyCostUSD
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// TableName returns the table name for the Diet schema.
func (Diet) TableName() string {
	return table.Name(module.Name, "diets")
}
