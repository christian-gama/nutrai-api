package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/query"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// FindPatientHandler represents the query handler for the FindPatient use case.
type FindPatientHandler = query.Handler[*FindPatientInput, *FindPatientOutput]

// findPatientHandlerImpl is the implementation of the FindPatient query handler.
type findPatientHandlerImpl struct {
	repo.Patient
}

// NewFindPatientHandler instantiates the FindPatient use case handler.
func NewFindPatientHandler(repo repo.Patient) FindPatientHandler {
	return &findPatientHandlerImpl{repo}
}

// Handle implements query.Handler.
func (q *findPatientHandlerImpl) Handle(ctx context.Context, input *FindPatientInput) (*FindPatientOutput, error) {
	patient, err := q.Patient.Find(ctx, repo.FindPatientInput{ID: input.ID}, "User")
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&FindPatientOutput{}, patient), nil
}
