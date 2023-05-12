package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/query"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// AllPatientsInput represents the input data for the AllPatients use case.
type AllPatientsHandler = query.Handler[*AllPatientsInput, *querying.PaginationOutput[*AllPatientsOutput]]

// allPatientsHandlerImpl is the implementation of the AllPatients use case handler.
type allPatientsHandlerImpl struct {
	repo.Patient
}

// NewAllPatientsHandler instantiates the AllPatients use case handler.
func NewAllPatientsHandler(repo repo.Patient) AllPatientsHandler {
	return &allPatientsHandlerImpl{repo}
}

// Handle implements query.Handler.
func (q *allPatientsHandlerImpl) Handle(ctx context.Context, input *AllPatientsInput) (*querying.PaginationOutput[*AllPatientsOutput], error) {
	pagination, err := q.Patient.All(ctx, repo.AllPatientsInput{
		Filterer:  input.Filter,
		Paginator: &input.Pagination,
		Sorter:    input.Sort,
		Preloader: input.Preload,
	})
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&querying.PaginationOutput[*AllPatientsOutput]{}, pagination), nil
}
