package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
)

// UpdatePatientHandler represents the UpdatePatient command.
type UpdatePatientHandler = command.Handler[*UpdatePatientInput]

// updatePatientHandlerImpl represents the implementation of the UpdatePatient command.
type updatePatientHandlerImpl struct {
	repo.Patient
}

// NewUpdatePatientHandler returns a new Update instance.
func NewUpdatePatientHandler(patientRepo repo.Patient) UpdatePatientHandler {
	if patientRepo == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	return &updatePatientHandlerImpl{patientRepo}
}

// Handle implements command.Handler.
func (c *updatePatientHandlerImpl) Handle(ctx context.Context, input *UpdatePatientInput) error {
	savedPatient, err := c.Find(ctx,
		repo.FindPatientInput{
			ID: input.ID,
		},
	)
	if err != nil {
		return err
	}

	allergies := make([]*patient.Allergy, len(input.Allergies))
	for i, allergy := range input.Allergies {
		allergies[i] = patient.NewAllergy().SetName(allergy.Name)
	}

	patient, err := savedPatient.
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
		SetAllergies(allergies).
		Validate()
	if err != nil {
		return err
	}

	return c.Update(ctx, repo.UpdatePatientInput{Patient: patient, ID: input.ID})
}
