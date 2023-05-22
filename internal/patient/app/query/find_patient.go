package query

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
)

// FindPatientInput is the input data of the patient query.
type FindPatientInput struct {
	ID coreValue.ID `uri:"id" faker:"uint" validate:"required,min=1"`
}

// FindPatientOutput is the output data of the patient query.
type FindPatientOutput struct {
	ID       coreValue.ID   `json:"id"`
	Age      value.Age      `json:"age"`
	HeightM  value.HeightM  `json:"heightM"`
	WeightKG value.WeightKG `json:"weightKG"`
	UserID   coreValue.ID   `json:"user,omitempty"`
	BMI      value.BMI      `json:"bmi"`
}
