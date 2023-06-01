package controller_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ChangePasswordSuite struct {
	suite.Suite
}

func TestChangePasswordSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ChangePasswordSuite))
}

func (s *ChangePasswordSuite) TestHandle() {
	type Mock struct {
		ChangePasswordHandler *cmdMock.Handler[*command.ChangePasswordInput]
	}

	type Sut struct {
		Sut   controller.ChangePassword
		Input *command.ChangePasswordInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.ChangePasswordInput()
		mock := &Mock{
			ChangePasswordHandler: cmdMock.NewHandler[*command.ChangePasswordInput](s.T()),
		}
		sut := controller.NewChangePassword(mock.ChangePasswordHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should change the current user's password", func() {
		sut := makeSut()

		sut.Mock.ChangePasswordHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.ChangePasswordHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("Password", func() {
		s.Run("should return bad request when empty", func() {
			sut := makeSut()

			sut.Input.Password = ""

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return bad request when longer than 32 chars", func() {
			sut := makeSut()

			sut.Input.Password = value.Password(strings.Repeat("a", 33))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})

		s.Run("should return bad request when shorter than 8 chars", func() {
			sut := makeSut()

			sut.Input.Password = value.Password(strings.Repeat("a", 7))

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when ChangePasswordHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.ChangePasswordHandler.
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
