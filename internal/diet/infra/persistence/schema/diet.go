package schema

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/user/infra/persistence/sql/schema"
)

// Diet is the database schema for diets.
type Diet struct {
	ID              uint `gorm:"primaryKey"`
	PatientID       uint
	Patient         *schema.Patient `gorm:"foreignKey:PatientID"`
	Description     string
	RestrictedFood  []string
	DurationInWeeks int16
	Goal            string
	MealPlan        string
	MonthlyCostUSD  float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// TableName returns the table name for the Diet schema.
func (d Diet) TableName() string {
	return "diets"
}
