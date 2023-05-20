package controller_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/api/controller"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	commandMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/command"
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
	type Mock struct {
		SavePatientHandler *commandMock.Handler[*command.SavePatientInput]
	}

	type Sut struct {
		Sut   controller.SavePatient
		Input *command.SavePatientInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.SavePatientInput()
		mock := &Mock{
			SavePatientHandler: commandMock.NewHandler[*command.SavePatientInput](s.T()),
		}
		sut := controller.NewSavePatient(mock.SavePatientHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should save a patient", func() {
		sut := makeSut()

		sut.Mock.SavePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.Mock.SavePatientHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("Age", func() {
		s.Run("should return error when less than 18", func() {
			sut := makeSut()

			sut.Input.Age = 17

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Age = 101

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("WeightKG", func() {
		s.Run("should return error when less than 30", func() {
			sut := makeSut()

			sut.Input.WeightKG = 29

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 600", func() {
			sut := makeSut()

			sut.Input.WeightKG = 601

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("HeightM", func() {
		s.Run("should return error when less than 1", func() {
			sut := makeSut()

			sut.Input.HeightM = 0.99

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 3", func() {
			sut := makeSut()

			sut.Input.HeightM = 3.01

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Name", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Name = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Name = value.Name(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 2", func() {
			sut := makeSut()

			sut.Input.Name = value.Name("a")

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Email", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Email = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when invalid", func() {
			sut := makeSut()

			sut.Input.Email = "invalid_email"

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Password", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Password = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 32", func() {
			sut := makeSut()

			sut.Input.Password = value.Password(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 8", func() {
			sut := makeSut()

			sut.Input.Password = value.Password("a")

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Age", func() {
		s.Run("should return error when less than 18", func() {
			sut := makeSut()

			sut.Input.Age = 17

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Age = 101

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("WeightKG", func() {
		s.Run("should return error when less than 30", func() {
			sut := makeSut()

			sut.Input.WeightKG = 29

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 600", func() {
			sut := makeSut()

			sut.Input.WeightKG = 601

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("HeightM", func() {
		s.Run("should return error when less than 1", func() {
			sut := makeSut()

			sut.Input.HeightM = 0.99

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 3", func() {
			sut := makeSut()

			sut.Input.HeightM = 3.01

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Name", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Name = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 100", func() {
			sut := makeSut()

			sut.Input.Name = value.Name(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 2", func() {
			sut := makeSut()

			sut.Input.Name = value.Name("a")

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Email", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Email = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when invalid", func() {
			sut := makeSut()

			sut.Input.Email = "invalid_email"

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("Password", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Password = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when greater than 32", func() {
			sut := makeSut()

			sut.Input.Password = value.Password(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 8", func() {
			sut := makeSut()

			sut.Input.Password = value.Password("a")

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when SavePatientHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.SavePatientHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
