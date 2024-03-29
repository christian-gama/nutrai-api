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

type LoginSuite struct {
	suite.Suite
}

func TestLoginSuite(t *testing.T) {
	suite.RunUnitTest(t, new(LoginSuite))
}

func (s *LoginSuite) TestHandle() {
	type Mock struct {
		LoginHandler *svcMock.Handler[*service.LoginInput, *service.LoginOutput]
	}

	type Sut struct {
		Sut   controller.Login
		Input *service.LoginInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.LoginInput()
		mocks := &Mock{
			LoginHandler: svcMock.NewHandler[*service.LoginInput, *service.LoginOutput](s.T()),
		}
		sut := controller.NewLogin(mocks.LoginHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: input}
	}

	s.Run("should return an access token and refresh token when login succeeds", func() {
		sut := makeSut()

		loginOutput := fake.LoginOutput()
		sut.Mock.LoginHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(loginOutput, nil)

		ctx, body := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Data: sut.Input,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		s.EqualValues(
			loginOutput.Access,
			body["access"],
			"should return access token",
		)
		s.EqualValues(
			loginOutput.Refresh,
			body["refresh"],
			"should return refresh token",
		)
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

	s.Run("panics when LoginHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.LoginHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
