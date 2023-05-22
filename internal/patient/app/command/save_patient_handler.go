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
	allergies := make([]*patient.Allergy, len(input.Allergies))
	for i, allergy := range input.Allergies {
		allergies[i] = patient.NewAllergy().SetName(allergy.Name)
	}

	patient, err := patient.NewPatient().
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
		SetID(input.User.ID).
		SetAllergies(allergies).
		Validate()
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SavePatientInput{Patient: patient}))
}
