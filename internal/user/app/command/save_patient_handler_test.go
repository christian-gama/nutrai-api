package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	fakePatient "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	userServiceMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/app/service"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SavePatientHandlerSuite struct {
	suite.Suite
}

func TestSaveHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SavePatientHandlerSuite))
}

func (s *SavePatientHandlerSuite) TestSaveHandler() {
	type Mocks struct {
		HashPasswordHandler *userServiceMock.HashPasswordHandler
		PatientRepo         *userRepoMock.Patient
	}

	type Sut struct {
		Sut   command.SavePatientHandler
		Ctx   context.Context
		Input *command.SavePatientInput
		Mocks *Mocks
	}

	makeSut := func() Sut {
		hashPasswordHandler := userServiceMock.NewHashPasswordHandler(s.T())
		patientRepo := userRepoMock.NewPatient(s.T())

		return Sut{
			Sut:   command.NewSavePatientHandler(patientRepo, hashPasswordHandler),
			Ctx:   context.Background(),
			Input: fake.SavePatientInput(),
			Mocks: &Mocks{hashPasswordHandler, patientRepo},
		}
	}

	s.Run("Should return nil when saving patient succeeds", func() {
		sut := makeSut()

		sut.Mocks.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mocks.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should call hashPasswordHandler.Handle with the password", func() {
		sut := makeSut()

		password := sut.Input.User.Password
		sut.Mocks.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mocks.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mocks.HashPasswordHandler.AssertCalled(s.T(), "Handle", sut.Ctx, &service.HashPasswordInput{Password: password})
	})

	s.Run("Should return error when hashing password fails", func() {
		sut := makeSut()

		sut.Mocks.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mocks.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)
		sut.Input.User.Email = ""

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when saving patient fails", func() {
		sut := makeSut()

		sut.Mocks.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mocks.PatientRepo.On("Save", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
