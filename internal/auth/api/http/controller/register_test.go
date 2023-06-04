package controller_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/service"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	svcMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/service"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RegisterSuite struct {
	suite.Suite
}

func TestRegisterSuite(t *testing.T) {
	suite.RunUnitTest(t, new(RegisterSuite))
}

func (s *RegisterSuite) TestHandle() {
	type Mock struct {
		RegisterHandler *svcMock.Handler[*service.RegisterInput, *service.RegisterOutput]
	}

	type Sut struct {
		Sut   controller.Register
		Input *service.RegisterInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.RegisterInput()
		mock := &Mock{
			RegisterHandler: svcMock.NewHandler[*service.RegisterInput, *service.RegisterOutput](
				s.T(),
			),
		}
		sut := controller.NewRegister(mock.RegisterHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should register a patient", func() {
		sut := makeSut()

		registerOutput := fake.RegisterOutput()
		sut.Mock.RegisterHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(registerOutput, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.Mock.RegisterHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

			sut.Input.Name = userValue.Name(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 2", func() {
			sut := makeSut()

			sut.Input.Name = userValue.Name("a")

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

			sut.Input.Password = userValue.Password(strings.Repeat("a", 101))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when less than 8", func() {
			sut := makeSut()

			sut.Input.Password = userValue.Password("a")

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data: sut.Input,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when RegisterHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.RegisterHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
