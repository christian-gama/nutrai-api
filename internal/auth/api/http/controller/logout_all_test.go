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

type LogoutAllSuite struct {
	suite.Suite
}

func TestLogoutAllSuite(t *testing.T) {
	suite.RunUnitTest(t, new(LogoutAllSuite))
}

func (s *LogoutAllSuite) TestHandle() {
	type Mock struct {
		LogoutAllHandler *cmdMock.Handler[*command.LogoutAllInput]
	}

	type Sut struct {
		Sut         controller.LogoutAll
		Input       *command.LogoutAllInput
		Mock        *Mock
		CurrentUser *user.User
	}

	makeSut := func() *Sut {
		input := cmdFake.LogoutAllInput()
		mocks := &Mock{
			LogoutAllHandler: cmdMock.NewHandler[*command.LogoutAllInput](s.T()),
		}
		sut := controller.NewLogoutAll(mocks.LogoutAllHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: input, CurrentUser: input.User}
	}

	s.Run("should logout all the current user's refresh tokens", func() {
		sut := makeSut()

		sut.Mock.LogoutAllHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.CurrentUser,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
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

	s.Run("panics when LogoutAllHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.LogoutAllHandler.
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
