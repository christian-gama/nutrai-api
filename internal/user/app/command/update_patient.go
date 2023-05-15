package command

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// UpdatePatientInput represents the input data for the UpdatePatient command.
type UpdatePatientInput struct {
	ID coreValue.ID `url:"id" validate:"required,min=1" faker:"boundary_start=1,boundary_end=100"`

	Age      value.Age        `json:"age" validate:"required,number,min=18,max=100" faker:"boundary_start=18,boundary_end=100"`
	HeightM  value.HeightM    `json:"heightM" validate:"required,number,min=1,max=3" faker:"boundary_start=1,boundary_end=2"`
	WeightKG value.WeightKG   `json:"weightKG" validate:"required,number,min=30,max=600" faker:"boundary_start=30,boundary_end=100"`
	User     *UpdateUserInput `json:"user" validate:"required,dive" faker:"-"`
}
