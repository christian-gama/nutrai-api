package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindPatientHandlerSuite struct {
	suite.Suite
}

func TestFindPatientHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindPatientHandlerSuite))
}

func (s *FindPatientHandlerSuite) TestPatientHandler() {
	type Mock struct {
		PatientRepo *mocks.Patient
	}

	type Sut struct {
		Sut   query.FindPatientHandler
		Ctx   context.Context
		Input *query.FindPatientInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PatientRepo: mocks.NewPatient(s.T()),
		}

		input := queryFake.FindPatientInput()

		sut := query.NewFindPatientHandler(mock.PatientRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return a PatientOutput", func() {
		sut := makeSut()

		patient := fake.Patient()
		patient.ID = sut.Input.ID
		sut.Mock.PatientRepo.On("Find", sut.Ctx, mock.Anything).Return(patient, nil)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Equal(sut.Input.ID, output.ID, "ID should be the same")
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.On("Find", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorIs(assert.AnError, err)
	})
}
