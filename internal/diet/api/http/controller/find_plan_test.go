package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/diet/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	userFake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/query"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindPlanSuite struct {
	suite.Suite
}

func TestFindPlanSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindPlanSuite))
}

func (s *FindPlanSuite) TestHandle() {
	type Mock struct {
		FindPlanHandler *mocks.Handler[*query.FindPlanInput, *query.FindPlanOutput]
	}

	type Sut struct {
		Sut         controller.FindPlan
		Input       query.FindPlanInput
		CurrentUser *user.User
		Mock        *Mock
	}

	makeSut := func() *Sut {
		input := fake.FindPlanInput()

		mock := &Mock{
			FindPlanHandler: mocks.NewHandler[*query.FindPlanInput, *query.FindPlanOutput](
				s.T(),
			),
		}

		currentUser := userFake.User()

		sut := controller.NewFindPlan(mock.FindPlanHandler)

		return &Sut{Sut: sut, Mock: mock, Input: *input, CurrentUser: currentUser}
	}

	s.Run("should find one plan successfuly", func() {
		sut := makeSut()

		findPlanOutput := fake.FindPlanOutput()
		findPlanOutput.ID = sut.Input.ID
		sut.Mock.FindPlanHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(findPlanOutput, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
			CurrentUser: sut.CurrentUser,
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.FindPlanHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("ID", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.ID = 0

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.CurrentUser,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when FindPlanHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.FindPlanHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params:      []string{fmt.Sprintf("%v", sut.Input.ID)},
				CurrentUser: sut.CurrentUser,
			})
		})
	})
}
