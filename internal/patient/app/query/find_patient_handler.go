package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/query"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// FindPatientHandler represents the query handler for the FindPatient use case.
type FindPatientHandler = query.Handler[*FindPatientInput, *FindPatientOutput]

// findPatientHandlerImpl is the implementation of the FindPatient query handler.
type findPatientHandlerImpl struct {
	repo.Patient
}

// NewFindPatientHandler instantiates the FindPatient use case handler.
func NewFindPatientHandler(patientRepo repo.Patient) FindPatientHandler {
	if patientRepo == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	return &findPatientHandlerImpl{patientRepo}
}

// Handle implements query.Handler.
func (q *findPatientHandlerImpl) Handle(
	ctx context.Context,
	input *FindPatientInput,
) (*FindPatientOutput, error) {
	p, err := q.Patient.Find(
		ctx,
		repo.FindPatientInput{
			ID:        input.ID,
			Preloader: input.Preload,
		},
	)
	if err != nil {
		return nil, err
	}

	return &FindPatientOutput{
		ID:       p.ID,
		Age:      p.Age,
		HeightM:  p.HeightM,
		WeightKG: p.WeightKG,
		BMI:      p.BMI,
		Allergies: slice.
			Map(p.Allergies, func(a *patient.Allergy) *FindPatientAllergiesOutput {
				return &FindPatientAllergiesOutput{
					ID:        a.ID,
					Name:      a.Name,
					PatientID: a.PatientID,
				}
			}).
			Build(),
	}, nil
}
