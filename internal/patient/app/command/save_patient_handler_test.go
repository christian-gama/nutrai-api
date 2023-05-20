package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/command"
	fakePatient "github.com/christian-gama/nutrai-api/testutils/fake/patient/domain/model/patient"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/patient/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SavePatientHandlerSuite struct {
	suite.Suite
}

func TestSavePatientHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SavePatientHandlerSuite))
}

func (s *SavePatientHandlerSuite) TestSaveHandler() {
	type Mock struct {
		PatientRepo *userRepoMock.Patient
	}

	type Sut struct {
		Sut   command.SavePatientHandler
		Ctx   context.Context
		Input *command.SavePatientInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PatientRepo: userRepoMock.NewPatient(s.T()),
		}

		input := fake.SavePatientInput()

		sut := command.NewSavePatientHandler(mock.PatientRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return nil when saving patient succeeds", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Input.Age = 0

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when saving patient fails", func() {
		sut := makeSut()

		sut.Mock.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
