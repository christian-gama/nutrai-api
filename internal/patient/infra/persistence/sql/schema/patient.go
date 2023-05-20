package schema

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
)

// Patient is the database schema for a patient.
type Patient struct {
	ID       coreValue.ID `gorm:"primaryKey"`
	WeightKG value.WeightKG
	HeightM  value.HeightM
	Age      value.Age
	BMI      value.BMI `gorm:"->"`
}

// TableName returns the table name for the Patient schema.
func (p Patient) TableName() string {
	return "patients"
}
