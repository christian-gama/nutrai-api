package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	qryMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/query"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AllPlansSuite struct {
	suite.Suite
}

func TestAllPlansSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllPlansSuite))
}

func (s *AllPlansSuite) TestHandle() {
	type Mock struct {
		AllPlansHandler *qryMock.Handler[*query.AllPlansInput, *queryer.PaginationOutput[*query.AllPlansOutput]]
	}

	type Sut struct {
		Sut   controller.AllPlans
		Input query.AllPlansInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.AllPlansInput()
		mocks := &Mock{
			AllPlansHandler: qryMock.NewHandler[*query.AllPlansInput, *queryer.PaginationOutput[*query.AllPlansOutput]](
				s.T(),
			),
		}
		sut := controller.NewAllPlans(mocks.AllPlansHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: *input}
	}

	s.Run("should fetch all patients successfuly", func() {
		sut := makeSut()

		sut.Mock.AllPlansHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&queryer.PaginationOutput[*query.AllPlansOutput]{
				Total: 1,
				Results: []*query.AllPlansOutput{
					&query.FindPlanOutput{},
				},
			}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.AllPlansHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("should fetch all patients successfuly using queries", func() {
		sut := makeSut()

		sut.Mock.AllPlansHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&queryer.PaginationOutput[*query.AllPlansOutput]{
				Total: 1,
				Results: []*query.AllPlansOutput{
					&query.FindPlanOutput{},
				},
			}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Queries: gintest.BuildScopeQuery(
				gintest.FilterOption(sut.Input.Filter.
					Add("age", "eq", 123).
					Add("weightKG", "eq", 123).
					Add("heightM", "eq", 123).
					Slice()),

				gintest.SortOption(sut.Input.Sort.
					Add("age", false).
					Add("weightKG", false).
					Add("heightM", false).
					Slice()),
			),
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.AllPlansHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("invalid Filter", func() {
		sut := makeSut()

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Queries: gintest.BuildScopeQuery(
				gintest.FilterOption(
					sut.Input.Filter.Add("invalid", "invalid", "invalid").Slice(),
				),
			),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("invalid Sort", func() {
		sut := makeSut()

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Queries: gintest.BuildScopeQuery(
				gintest.SortOption(sut.Input.Sort.Add("invalid", false).Slice()),
			),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when AllPlansHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.AllPlansHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
