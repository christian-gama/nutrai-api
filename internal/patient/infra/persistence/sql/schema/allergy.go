package schema

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/core/infra/table"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/module"
)

// Allergy is the database schema for a patient.
type Allergy struct {
	ID        coreValue.ID `gorm:"primaryKey"`
	PatientID coreValue.ID
	Name      value.Allergy
}

// TableName returns the table name for the Allergy schema.
func (p Allergy) TableName() string {
	return table.Name(module.Module, "allergies")
}
