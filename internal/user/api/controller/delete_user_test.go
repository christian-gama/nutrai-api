package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/api/controller"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	commandMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteUserSuite struct {
	suite.Suite
}

func TestDeleteUserSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteUserSuite))
}

func (s *DeleteUserSuite) TestHandle() {
	type Mock struct {
		DeleteUserHandler *commandMock.Handler[*command.DeleteUserInput]
	}

	type Sut struct {
		Sut   controller.DeleteUser
		Input *command.DeleteUserInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.DeleteUserInput()
		mock := &Mock{DeleteUserHandler: commandMock.NewHandler[*command.DeleteUserInput](s.T())}
		sut := controller.NewDeleteUser(mock.DeleteUserHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should delete a user", func() {
		sut := makeSut()

		sut.Mock.DeleteUserHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:   sut.Input,
			Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.Mock.DeleteUserHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("ID", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.ID = 0

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:   sut.Input,
				Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when DeleteUserHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.DeleteUserHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:   sut.Input,
				Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
			})
		})
	})
}
