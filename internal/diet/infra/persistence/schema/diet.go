package schema

import "github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"

// Diet is the database schema for diets.
type Diet struct {
	ID              uint `gorm:"primaryKey"`
	PatientID       uint
	Patient         *patient.Patient `gorm:"foreignKey:PatientID"`
	Description     string
	RestrictedFood  []string
	DurationInWeeks int16
	Goal            string
	MealPlan        string
	MonthlyCostUSD  float64
}

// TableName returns the table name for the Diet schema.
func (d Diet) TableName() string {
	return "diets"
}
