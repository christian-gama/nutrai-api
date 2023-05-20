package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/query"
	"github.com/christian-gama/nutrai-api/internal/core/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
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
	patient, err := q.Patient.Find(
		ctx,
		repo.FindPatientInput{ID: input.ID, Preloader: input.Preload},
	)
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&FindPatientOutput{}, patient), nil
}
