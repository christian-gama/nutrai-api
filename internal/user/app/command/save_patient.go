package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"

// SavePatientInput represents the input data for the SavePatient command.
type SavePatientInput struct {
	Age      value.Age      `json:"age" faker:"boundary_start=18,boundary_end=100"`
	HeightM  value.HeightM  `json:"heightM" faker:"boundary_start=1,boundary_end=2"`
	WeightKG value.WeightKG `json:"weightKG" faker:"boundary_start=30,boundary_end=100"`
	User     *SaveUserInput `json:"user" faker:"-"`
}
