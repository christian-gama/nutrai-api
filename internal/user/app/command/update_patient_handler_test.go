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
	type Mock struct {
		PatientRepo *userRepoMock.Patient
	}

	type Sut struct {
		Sut   command.UpdatePatientHandler
		Ctx   context.Context
		Input *command.UpdatePatientInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PatientRepo: userRepoMock.NewPatient(s.T()),
		}

		input := fake.UpdatePatientInput()

		sut := command.NewUpdatePatientHandler(mock.PatientRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return nil when saving patient succeeds", func() {
		sut := makeSut()

		patient := fakePatient.Patient()
		sut.Mock.PatientRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(patient, nil)

		sut.Mock.PatientRepo.
			On("Update", sut.Ctx, mock.Anything).
			Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		sut.Input.Email = ""

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when saving user fails", func() {
		sut := makeSut()

		patient := fakePatient.Patient()
		sut.Mock.PatientRepo.
			On("Find", sut.Ctx, mock.Anything).
			Return(patient, nil)

		sut.Mock.PatientRepo.On("Update", sut.Ctx, mock.Anything).Return(assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
