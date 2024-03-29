package controller_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	"github.com/christian-gama/nutrai-api/internal/patient/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/patient/app/command"
	value "github.com/christian-gama/nutrai-api/internal/patient/domain/value/patient"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
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
		Sut                controller.SavePatient
		Input              *command.SavePatientInput
		SavePatientHandler *cmdMock.Handler[*command.SavePatientInput]
	}

	makeSut := func() *Sut {
		input := fake.SavePatientInput()
		savePatient := cmdMock.NewHandler[*command.SavePatientInput](s.T())
		sut := controller.NewSavePatient(savePatient)
		return &Sut{Sut: sut, SavePatientHandler: savePatient, Input: input}
	}

	s.Run("should save a patient", func() {
		sut := makeSut()

		sut.SavePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.SavePatientHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("Age", func() {
		s.Run("should return error when less than 18", func() {
			sut := makeSut()

			sut.Input.Age = 17

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Age = 101

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("WeightKG", func() {
		s.Run("should return error when less than 30", func() {
			sut := makeSut()

			sut.Input.WeightKG = 29

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 600", func() {
			sut := makeSut()

			sut.Input.WeightKG = 601

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("HeightM", func() {
		s.Run("should return error when less than 1", func() {
			sut := makeSut()

			sut.Input.HeightM = 0.99

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 3", func() {
			sut := makeSut()

			sut.Input.HeightM = 3.01

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Allergies", func() {
		s.Run("should return error when invalid", func() {
			sut := makeSut()

			sut.Input.Allergies = []value.Allergy{value.Allergy(strings.Repeat("a", 101))}

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("User", func() {
		s.Run("should panic when empty", func() {
			sut := makeSut()

			sut.Input.User = nil

			s.PanicsWithValue(ctxstore.ErrUserNotFound, func() {
				gintest.MustRequestWithBody(sut.Sut, gintest.Option{
					Data:        sut.Input,
					CurrentUser: sut.Input.User,
				})
			})
		})
	})

	s.Run("panics when SavePatientHandler.Handle returns error", func() {
		sut := makeSut()

		sut.SavePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})
		})
	})
}
