package command

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// UpdatePatientInput represents the input data for the UpdatePatient command.
type UpdatePatientInput struct {
	ID sharedvalue.ID `form:"id"`

	Age      value.Age        `json:"age" faker:"boundary_start=18,boundary_end=100"`
	HeightM  value.HeightM    `json:"heightM" faker:"boundary_start=1,boundary_end=2"`
	WeightKG value.WeightKG   `json:"weightKG" faker:"boundary_start=30,boundary_end=100"`
	User     *UpdateUserInput `json:"user" faker:"-"`
}
