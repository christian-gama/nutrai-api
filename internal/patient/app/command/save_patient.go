package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
)

// SavePatientInput represents the input data for the SavePatient command.
type SavePatientInput struct {
	User *user.User `ctx:"currentUser" json:"-"`

	Age       value.Age           `json:"age" validate:"required,number,min=18,max=100" faker:"boundary_start=18,boundary_end=100"`
	HeightM   value.HeightM       `json:"heightM" validate:"required,number,min=1,max=3" faker:"boundary_start=1,boundary_end=2"`
	WeightKG  value.WeightKG      `json:"weightKG" validate:"required,number,min=30,max=600" faker:"boundary_start=30,boundary_end=100"`
	Allergies []*SaveAllergyInput `json:"allergies" validate:"required,dive,required"`
}

type SaveAllergyInput struct {
	Name value.Allergy `json:"name" validate:"required,max=100" faker:"len=50"`
}
