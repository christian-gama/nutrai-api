package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	"github.com/christian-gama/nutrai-api/internal/patient/domain/model/patient"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AllPatientsHandlerSuite struct {
	suite.Suite
}

func TestAllPatientsHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllPatientsHandlerSuite))
}

func (s *AllPatientsHandlerSuite) TestPatientHandler() {
	type Mock struct {
		PatientRepo *mocks.Patient
	}

	type Sut struct {
		Sut   query.AllPatientsHandler
		Ctx   context.Context
		Input *query.AllPatientsInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PatientRepo: mocks.NewPatient(s.T()),
		}

		input := queryFake.AllPatientsInput()

		sut := query.NewAllPatientsHandler(mock.PatientRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return a PatientOutput", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.On("All", sut.Ctx, mock.Anything).
			Return(
				&queryer.PaginationOutput[*patient.Patient]{
					Results: []*patient.Patient{fake.Patient()},
					Total:   1,
				},
				nil,
			)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Len(output.Results, 1)
		s.Equal(1, output.Total)
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.On("All", sut.Ctx, mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorIs(assert.AnError, err)
	})
}
