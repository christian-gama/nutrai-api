package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/api/controller"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	commandMock "github.com/christian-gama/nutrai-api/testutils/mocks/shared/app/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SavePatientSuite struct {
	suite.Suite
}

func TestSavePatientSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SavePatientSuite))
}

func (s *SavePatientSuite) TestHandle() {
	type Sut struct {
		Sut         controller.SavePatient
		Input       *command.SavePatientInput
		SavePatient *commandMock.Handler[*command.SavePatientInput]
	}

	makeSut := func() *Sut {
		input := fake.SavePatientInput()
		savePatient := commandMock.NewHandler[*command.SavePatientInput](s.T())
		sut := controller.NewSavePatient(savePatient)
		return &Sut{Sut: sut, SavePatient: savePatient, Input: input}
	}

	s.Run("should save a patient", func() {
		sut := makeSut()

		sut.SavePatient.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.SavePatient.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid Age: it's required", func() {
		sut := makeSut()

		sut.Input.Age = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when SavePatient.Handle returns error", func() {
		sut := makeSut()

		sut.SavePatient.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
