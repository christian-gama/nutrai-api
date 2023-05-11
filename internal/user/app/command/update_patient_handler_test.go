package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	fakePatient "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdatePatientHandlerSuite struct {
	suite.Suite
}

func TestUpdateHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdatePatientHandlerSuite))
}

func (s *UpdatePatientHandlerSuite) TestUpdateHandler() {
	type Mocks struct {
		PatientRepo *userRepoMock.Patient
	}

	type Sut struct {
		Sut   command.UpdatePatientHandler
		Ctx   context.Context
		Input *command.UpdatePatientInput
		Mocks *Mocks
	}

	makeSut := func() Sut {
		patientRepo := userRepoMock.NewPatient(s.T())

		return Sut{
			Sut:   command.NewUpdatePatientHandler(patientRepo),
			Ctx:   context.Background(),
			Input: fake.UpdatePatientInput(),
			Mocks: &Mocks{patientRepo},
		}
	}

	s.Run("Should return nil when saving patient succeeds", func() {
		sut := makeSut()

		patient := fakePatient.Patient()
		sut.Mocks.PatientRepo.
			On("Find", sut.Ctx, mock.Anything, "User").
			Return(patient, nil)

		sut.Mocks.PatientRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mocks.PatientRepo.
			On("Find", sut.Ctx, mock.Anything, "User").
			Return(fakePatient.Patient(), nil)

		sut.Input.User.Email = ""

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when saving user fails", func() {
		sut := makeSut()

		patient := fakePatient.Patient()
		sut.Mocks.PatientRepo.
			On("Find", sut.Ctx, mock.Anything, "User").
			Return(patient, nil)

		sut.Mocks.PatientRepo.On("Update", sut.Ctx, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
