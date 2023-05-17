package command

import (
	patientValue "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
	userValue "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// SavePatientInput represents the input data for the SavePatient command.
type SavePatientInput struct {
	Age      patientValue.Age      `json:"age" validate:"required,number,min=18,max=100" faker:"boundary_start=18,boundary_end=100"`
	HeightM  patientValue.HeightM  `json:"heightM" validate:"required,number,min=1,max=3" faker:"boundary_start=1,boundary_end=2"`
	WeightKG patientValue.WeightKG `json:"weightKG" validate:"required,number,min=30,max=600" faker:"boundary_start=30,boundary_end=100"`
	Email    userValue.Email       `json:"email" validate:"required,email" faker:"email"`
	Password userValue.Password    `json:"password" validate:"required,alphanum,min=8,max=32" faker:"len=8"`
	Name     userValue.Name        `json:"name" validate:"required,min=2,max=100" faker:"name"`
}
