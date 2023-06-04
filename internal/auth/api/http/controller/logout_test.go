package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	cmdFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type LogoutSuite struct {
	suite.Suite
}

func TestLogoutSuite(t *testing.T) {
	suite.RunUnitTest(t, new(LogoutSuite))
}

func (s *LogoutSuite) TestHandle() {
	type Mock struct {
		LogoutHandler *cmdMock.Handler[*command.LogoutInput]
	}

	type Sut struct {
		Sut         controller.Logout
		Input       *command.LogoutInput
		Mock        *Mock
		CurrentUser *user.User
	}

	makeSut := func() *Sut {
		input := cmdFake.LogoutInput()
		mocks := &Mock{
			LogoutHandler: cmdMock.NewHandler[*command.LogoutInput](s.T()),
		}
		sut := controller.NewLogout(mocks.LogoutHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: input, CurrentUser: input.User}
	}

	s.Run("should logout the current user", func() {
		sut := makeSut()

		sut.Mock.LogoutHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.CurrentUser,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})

	s.Run("Refresh", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.Refresh = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.CurrentUser,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return error when invalid", func() {
			sut := makeSut()

			sut.Input.Refresh = "invalid_token"

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.CurrentUser,
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

	s.Run("panics when LogoutHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.LogoutHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.CurrentUser,
			})
		})
	})
}
