package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	fakePatient "github.com/christian-gama/nutrai-api/testutils/fake/user/domain/model/patient"
	serviceMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/service"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/user/domain/repo"
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
		HashPasswordHandler *serviceMock.Handler[*service.HashPasswordInput, *service.HashPasswordOutput]
		PatientRepo         *userRepoMock.Patient
	}

	type Sut struct {
		Sut   command.SavePatientHandler
		Ctx   context.Context
		Input *command.SavePatientInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			HashPasswordHandler: serviceMock.NewHandler[*service.HashPasswordInput, *service.HashPasswordOutput](
				s.T(),
			),
			PatientRepo: userRepoMock.NewPatient(s.T()),
		}

		input := fake.SavePatientInput()

		sut := command.NewSavePatientHandler(mock.PatientRepo, mock.HashPasswordHandler)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should return nil when saving patient succeeds", func() {
		sut := makeSut()

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should call hashPasswordHandler.Handle with the password", func() {
		sut := makeSut()

		password := sut.Input.Password
		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.PatientRepo.
			On("Save", sut.Ctx, mock.Anything).
			Return(fakePatient.Patient(), nil)

		_ = sut.Sut.Handle(sut.Ctx, sut.Input)

		sut.Mock.HashPasswordHandler.AssertCalled(
			s.T(),
			"Handle",
			sut.Ctx,
			&service.HashPasswordInput{Password: password},
		)
	})

	s.Run("Should return error when hashing password fails", func() {
		sut := makeSut()

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			// Must return a pointer to output because of generics.
			Return(&service.HashPasswordOutput{}, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})

	s.Run("Should return error when converting input to model fails", func() {
		sut := makeSut()

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)
		sut.Input.Email = ""

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorAsRequired(err)
	})

	s.Run("Should return error when saving patient fails", func() {
		sut := makeSut()

		sut.Mock.HashPasswordHandler.
			On("Handle", sut.Ctx, mock.Anything).
			Return(&service.HashPasswordOutput{Password: "hashed"}, nil)

		sut.Mock.PatientRepo.On("Save", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
