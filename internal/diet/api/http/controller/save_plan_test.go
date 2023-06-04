package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/diet/app/service"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/service"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	svcMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/service"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SavePlanSuite struct {
	suite.Suite
}

func TestSavePlanSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SavePlanSuite))
}

func (s *SavePlanSuite) TestHandle() {
	type Sut struct {
		Sut             controller.SavePlan
		Input           *service.SavePlanInput
		SavePlanHandler *svcMock.Handler[*service.SavePlanInput, *service.SavePlanOutput]
	}

	makeSut := func() *Sut {
		input := fake.SavePlanInput()
		savePlan := svcMock.NewHandler[*service.SavePlanInput, *service.SavePlanOutput](s.T())
		sut := controller.NewSavePlan(savePlan)
		return &Sut{Sut: sut, SavePlanHandler: savePlan, Input: input}
	}

	s.Run("should save a plan", func() {
		sut := makeSut()

		sut.SavePlanHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(fake.SavePlanOutput(), nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Data:        sut.Input,
			CurrentUser: sut.Input.User,
		})

		s.Equal(http.StatusCreated, ctx.Writer.Status())
		sut.SavePlanHandler.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("DietID", func() {
		s.Run("should return error when less than 1", func() {
			sut := makeSut()

			sut.Input.DietID = 0

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when SavePlanHandler.Handle returns error", func() {
		sut := makeSut()

		sut.SavePlanHandler.
			On("Handle", mock.Anything, sut.Input).
			Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data:        sut.Input,
				CurrentUser: sut.Input.User,
			})
		})
	})

	s.Run("panics when user is not in context", func() {
		sut := makeSut()

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Data: sut.Input,
			})
		})
	})
}
