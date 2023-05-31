package controller_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
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

type UpdatePatientSuite struct {
	suite.Suite
}

func TestUpdatePatientSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdatePatientSuite))
}

func (s *UpdatePatientSuite) TestHandle() {
	type Mock struct {
		UpdatePatientHandler *cmdMock.Handler[*command.UpdatePatientInput]
	}

	type Sut struct {
		Sut   controller.UpdatePatient
		Input *command.UpdatePatientInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.UpdatePatientInput()

		mock := &Mock{
			UpdatePatientHandler: cmdMock.NewHandler[*command.UpdatePatientInput](s.T()),
		}

		sut := controller.NewUpdatePatient(mock.UpdatePatientHandler)

		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should update a patient", func() {
		sut := makeSut()

		sut.Mock.UpdatePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.UpdatePatientHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("ID", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.ID = 0

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Age", func() {
		s.Run("should return error when less than 18", func() {
			sut := makeSut()

			sut.Input.Age = 17

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Age = 101

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
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
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 600", func() {
			sut := makeSut()

			sut.Input.WeightKG = 601

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
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
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 3", func() {
			sut := makeSut()

			sut.Input.HeightM = 3.01

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
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
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("should panic when user is not in context", func() {
		sut := makeSut()

		s.PanicsWithValue(store.ErrUserNotFound, func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: nil,
			})
		})
	})

	s.Run("panics when UpdatePatientHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.UpdatePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})
		})
	})
}
