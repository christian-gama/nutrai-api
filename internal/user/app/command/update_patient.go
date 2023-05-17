package command

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	patientValue "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
	userValue "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// UpdatePatientInput represents the input data for the UpdatePatient command.
type UpdatePatientInput struct {
	ID coreValue.ID `uri:"id" validate:"required,min=1" faker:"boundary_start=1,boundary_end=100"`

	Age      patientValue.Age      `json:"age" validate:"required,number,min=18,max=100" faker:"boundary_start=18,boundary_end=100"`
	HeightM  patientValue.HeightM  `json:"heightM" validate:"required,number,min=1,max=3" faker:"boundary_start=1,boundary_end=2"`
	WeightKG patientValue.WeightKG `json:"weightKG" validate:"required,number,min=30,max=600" faker:"boundary_start=30,boundary_end=100"`
	Email    userValue.Email       `json:"email" validate:"required,email" faker:"email"`
	Name     userValue.Name        `json:"name" validate:"required,min=2,max=100" faker:"name"`
}
