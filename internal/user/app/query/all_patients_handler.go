package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/shared/app/query"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/convert"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// AllPatientsInput represents the input data for the AllPatients use case.
type AllPatientsHandler = query.Handler[*AllPatientsInput, *queryer.PaginationOutput[*AllPatientsOutput]]

// allPatientsHandlerImpl is the implementation of the AllPatients use case handler.
type allPatientsHandlerImpl struct {
	repo.Patient
}

// NewAllPatientsHandler instantiates the AllPatients use case handler.
func NewAllPatientsHandler(r repo.Patient) AllPatientsHandler {
	if r == nil {
		panic(errors.New("repo.Patient cannot be nil"))
	}

	return &allPatientsHandlerImpl{r}
}

// Handle implements query.Handler.
func (q *allPatientsHandlerImpl) Handle(ctx context.Context, input *AllPatientsInput) (*queryer.PaginationOutput[*AllPatientsOutput], error) {
	pagination, err := q.Patient.All(ctx, repo.AllPatientsInput{
		Filterer:  input.Filter,
		Paginator: &input.Pagination,
		Sorter:    input.Sort,
		Preloader: input.Preload,
	})
	if err != nil {
		return nil, err
	}

	return convert.FromModel(&queryer.PaginationOutput[*AllPatientsOutput]{}, pagination), nil
}
