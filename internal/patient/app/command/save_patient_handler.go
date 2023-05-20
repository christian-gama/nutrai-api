package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
)

// SavePatientHandler represents the SavePatient command.
type SavePatientHandler = command.Handler[*SavePatientInput]

// savePatientHandlerImpl represents the implementation of the SavePatient command.
type savePatientHandlerImpl struct {
	repo.Patient
}

// NewSavePatientHandler returns a new Save instance.
func NewSavePatientHandler(patientRepo repo.Patient) SavePatientHandler {
	if patientRepo == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	return &savePatientHandlerImpl{patientRepo}
}

// Handle implements command.Handler.
func (c *savePatientHandlerImpl) Handle(ctx context.Context, input *SavePatientInput) error {
	patient, err := patient.New().
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
		SetUserID(0).
		Build()
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SavePatientInput{Patient: patient}))
}
