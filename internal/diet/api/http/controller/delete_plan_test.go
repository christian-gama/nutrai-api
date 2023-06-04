package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/ctxstore"
	"github.com/christian-gama/nutrai-api/internal/diet/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/command"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	cmdMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/command"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeletePlanSuite struct {
	suite.Suite
}

func TestDeletePlanSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeletePlanSuite))
}

func (s *DeletePlanSuite) TestHandle() {
	type Mock struct {
		DeletePlanHandler *cmdMock.Handler[*command.DeletePlanInput]
	}

	type Sut struct {
		Sut   controller.DeletePlan
		Input *command.DeletePlanInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.DeletePlanInput()
		mock := &Mock{DeletePlanHandler: cmdMock.NewHandler[*command.DeletePlanInput](s.T())}
		sut := controller.NewDeletePlan(mock.DeletePlanHandler)
		return &Sut{Sut: sut, Mock: mock, Input: input}
	}

	s.Run("should delete the current user", func() {
		sut := makeSut()

		sut.Mock.DeletePlanHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			Params:      []string{fmt.Sprintf("%d", sut.Input.ID)},
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.Mock.DeletePlanHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("User", func() {
		s.Run("should panic when empty", func() {
			sut := makeSut()

			sut.Input.User = nil

			s.PanicsWithValue(ctxstore.ErrUserNotFound, func() {
				gintest.MustRequestWithBody(sut.Sut, gintest.Option{
					Data:        sut.Input,
					Params:      []string{fmt.Sprintf("%d", sut.Input.ID)},
					CurrentUser: sut.Input.User,
				})
			})
		})
	})

	s.Run("panics when DeletePlanHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.DeletePlanHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				Params:      []string{fmt.Sprintf("%d", sut.Input.ID)},
				CurrentUser: sut.Input.User,
			})
		})
	})
}
