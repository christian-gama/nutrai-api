package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
)

// UpdatePatientInput represents the input data for the UpdatePatient command.
type UpdatePatientInput struct {
	User *user.User `ctx:"current_user" json:"-"`

	Age       value.Age       `json:"age" validate:"required,patient_age" faker:"boundary_start=18,boundary_end=100"`
	HeightM   value.HeightM   `json:"heightM" validate:"required,patient_height_m" faker:"boundary_start=1,boundary_end=2"`
	WeightKG  value.WeightKG  `json:"weightKG" validate:"required,patient_weight_kg" faker:"boundary_start=30,boundary_end=100"`
	Allergies []value.Allergy `json:"allergies" validate:"required,unique,dive,patient_allergy_name" faker:"-"`
}
