package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/user/app/query"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/patient"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
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
	type Mocks struct {
		Repo *mocks.Patient
	}

	type Sut struct {
		Sut   query.AllPatientsHandler
		Ctx   context.Context
		Input *query.AllPatientsInput
		Mocks *Mocks
	}

	makeSut := func() Sut {
		patientRepo := mocks.NewPatient(s.T())

		return Sut{
			Sut:   query.NewAllPatientsHandler(patientRepo),
			Ctx:   context.Background(),
			Input: queryFake.AllPatientsInput(),
			Mocks: &Mocks{
				Repo: patientRepo,
			},
		}
	}

	s.Run("Should return a PatientOutput", func() {
		sut := makeSut()

		sut.Mocks.Repo.On("All", sut.Ctx, mock.Anything, "User").Return(&querying.PaginationOutput[*patient.Patient]{
			Results: []*patient.Patient{fake.Patient()},
			Total:   1,
		}, nil)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Len(output.Results, 1)
		s.Equal(1, output.Total)
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mocks.Repo.On("All", sut.Ctx, mock.Anything, "User").Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorIs(assert.AnError, err)
	})
}
