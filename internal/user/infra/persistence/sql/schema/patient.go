package schema

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// Patient is the database schema for a patient.
type Patient struct {
	ID       sharedvalue.ID `gorm:"primaryKey"`
	UserID   sharedvalue.ID
	User     *User `gorm:"foreignKey:UserID"`
	WeightKG value.WeightKG
	HeightM  value.HeightM
	Age      value.Age
	BMI      value.BMI `gorm:"->"`
}

// TableName returns the table name for the Patient schema.
func (p Patient) TableName() string {
	return "patients"
}
