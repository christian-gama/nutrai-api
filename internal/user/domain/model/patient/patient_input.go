package patient

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/patient"
)

// PatientInput is the input to create a new Patient.
type PatientInput struct {
	ID       sharedvalue.ID
	UserID   sharedvalue.ID
	WeightKG value.WeightKG
	HeightM  value.HeightM
	Age      value.Age
}
