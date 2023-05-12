package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/query"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
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
	type Mocks struct {
		Repo *mocks.Patient
	}

	type Sut struct {
		Sut   query.FindPatientHandler
		Ctx   context.Context
		Input *query.FindPatientInput
		Mocks *Mocks
	}

	makeSut := func() Sut {
		patientRepo := mocks.NewPatient(s.T())

		return Sut{
			Sut:   query.NewFindPatientHandler(patientRepo),
			Ctx:   context.Background(),
			Input: queryFake.FindPatientInput(),
			Mocks: &Mocks{
				Repo: patientRepo,
			},
		}
	}

	s.Run("Should return a PatientOutput", func() {
		sut := makeSut()

		patient := fake.Patient()
		patient.ID = sut.Input.ID
		sut.Mocks.Repo.On("Find", sut.Ctx, mock.Anything).Return(patient, nil)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.NoError(err)
		s.Require().NotNil(output)
		s.Equal(sut.Input.ID, output.ID, "ID should be the same")
	})

	s.Run("Should return an error when the repository fails", func() {
		sut := makeSut()

		sut.Mocks.Repo.On("Find", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Error(err)
		s.Nil(output)
		s.ErrorIs(assert.AnError, err)
	})
}
