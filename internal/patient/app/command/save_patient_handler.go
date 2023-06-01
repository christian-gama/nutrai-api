package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// SavePatientHandler represents the SavePatient command.
type SavePatientHandler = command.Handler[*SavePatientInput]

// savePatientHandlerImpl represents the implementation of the SavePatient command.
type savePatientHandlerImpl struct {
	repo.Patient
}

// NewSavePatientHandler returns a new Save instance.
func NewSavePatientHandler(patientRepo repo.Patient) SavePatientHandler {
	errutil.MustBeNotEmpty("repo.Patient", patientRepo)

	return &savePatientHandlerImpl{patientRepo}
}

// Handle implements command.Handler.
func (c *savePatientHandlerImpl) Handle(ctx context.Context, input *SavePatientInput) error {
	patient, err := patient.NewPatient().
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
		SetID(input.User.ID).
		SetAllergies(
			slice.
				Map(input.Allergies, func(allergy value.Allergy) *patient.Allergy {
					return patient.NewAllergy().SetName(allergy)
				}).
				Build()...,
		).
		Validate()
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SavePatientInput{Patient: patient}))
}
