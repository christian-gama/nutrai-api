package controller_test

import (
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/patient/api/controller"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/query"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/query"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AllPatientsSuite struct {
	suite.Suite
}

func TestAllPatientsSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllPatientsSuite))
}

func (s *AllPatientsSuite) TestHandle() {
	type Mock struct {
		AllPatientsHandler *mocks.Handler[*query.AllPatientsInput, *queryer.PaginationOutput[*query.AllPatientsOutput]]
	}

	type Sut struct {
		Sut   controller.AllPatients
		Input query.AllPatientsInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.AllPatientsInput()
		mocks := &Mock{
			AllPatientsHandler: mocks.NewHandler[*query.AllPatientsInput, *queryer.PaginationOutput[*query.AllPatientsOutput]](
				s.T(),
			),
		}
		sut := controller.NewAllPatients(mocks.AllPatientsHandler)
		return &Sut{Sut: sut, Mock: mocks, Input: *input}
	}

	s.Run("should fetch all patients successfuly", func() {
		sut := makeSut()

		sut.Mock.AllPatientsHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&queryer.PaginationOutput[*query.AllPatientsOutput]{
				Total: 1,
				Results: []*query.AllPatientsOutput{
					&query.FindPatientOutput{},
				},
			}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.AllPatientsHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("should fetch all patients successfuly using queries", func() {
		sut := makeSut()

		sut.Mock.AllPatientsHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&queryer.PaginationOutput[*query.AllPatientsOutput]{
				Total: 1,
				Results: []*query.AllPatientsOutput{
					&query.FindPatientOutput{},
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
		sut.Mock.AllPatientsHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
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

	s.Run("panics when AllPatientsHandler.Handle returns error", func() {
		sut := makeSut()

		sut.Mock.AllPatientsHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{})
		})
	})
}
