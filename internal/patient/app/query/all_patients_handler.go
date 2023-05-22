package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/query"
	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/repo"
)

// AllPatientsInput represents the input data for the AllPatients use case.
type AllPatientsHandler = query.Handler[*AllPatientsInput, *queryer.PaginationOutput[*AllPatientsOutput]]

// allPatientsHandlerImpl is the implementation of the AllPatients use case handler.
type allPatientsHandlerImpl struct {
	repo.Patient
}

// NewAllPatientsHandler instantiates the AllPatients use case handler.
func NewAllPatientsHandler(patientRepo repo.Patient) AllPatientsHandler {
	if patientRepo == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	return &allPatientsHandlerImpl{patientRepo}
}

// Handle implements query.Handler.
func (q *allPatientsHandlerImpl) Handle(
	ctx context.Context,
	input *AllPatientsInput,
) (*queryer.PaginationOutput[*AllPatientsOutput], error) {
	pagination, err := q.Patient.All(ctx, repo.AllPatientsInput{
		Filterer:  input.Filter,
		Paginator: &input.Pagination,
		Sorter:    input.Sort,
	})
	if err != nil {
		return nil, err
	}

	output := &queryer.PaginationOutput[*AllPatientsOutput]{
		Total:   pagination.Total,
		Results: []*FindPatientOutput{},
	}

	for _, patient := range pagination.Results {
		output.Results = append(output.Results, &FindPatientOutput{
			ID:       patient.ID,
			Age:      patient.Age,
			HeightM:  patient.HeightM,
			WeightKG: patient.WeightKG,
			UserID:   patient.UserID,
			BMI:      patient.BMI,
		})
	}

	return output, nil
}
