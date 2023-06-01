package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteMeSuite struct {
	suite.Suite
}

func TestDeleteMeSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteMeSuite))
}

func (s *DeleteMeSuite) TestHandle() {
	type Mock struct {
		DeleteMeHandler *cmdMock.Handler[*command.DeleteMeInput]
	}

	type Sut struct {
		Sut   controller.DeleteMe
		Input *command.DeleteMeInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.DeleteMeInput()
		mock := &Mock{DeleteMeHandler: cmdMock.NewHandler[*command.DeleteMeInput](s.T())}
		sut := controller.NewDeleteMe(mock.DeleteMeHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should delete the current user", func() {
		sut := makeSut()

		sut.Mock.DeleteMeHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.Mock.DeleteMeHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
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

	s.Run("panics when DeleteMeHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.DeleteMeHandler.
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
