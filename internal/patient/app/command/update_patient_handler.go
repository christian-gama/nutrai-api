package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// UpdatePatientHandler represents the UpdatePatient command.
type UpdatePatientHandler = command.Handler[*UpdatePatientInput]

// updatePatientHandlerImpl represents the implementation of the UpdatePatient command.
type updatePatientHandlerImpl struct {
	repo.Patient
}

// NewUpdatePatientHandler returns a new Update instance.
func NewUpdatePatientHandler(patientRepo repo.Patient) UpdatePatientHandler {
	errutil.MustBeNotEmpty("repo.Patient", patientRepo)

	return &updatePatientHandlerImpl{patientRepo}
}

// Handle implements command.Handler.
func (c *updatePatientHandlerImpl) Handle(ctx context.Context, input *UpdatePatientInput) error {
	savedPatient, err := c.Find(ctx,
		repo.FindPatientInput{
			ID:        input.User.ID,
			Preloader: querying.AddPreload("Allergies"),
		},
	)
	if err != nil {
		return err
	}

	patient, err := savedPatient.
		SetAge(input.Age).
		SetHeightM(input.HeightM).
		SetWeightKG(input.WeightKG).
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

	return c.Update(ctx, repo.UpdatePatientInput{Patient: patient, ID: input.User.ID})
}
