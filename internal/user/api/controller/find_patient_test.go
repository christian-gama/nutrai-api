package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/api/controller"
	"github.com/christian-gama/nutrai-api/internal/user/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/user/app/query"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/app/query"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindPatientSuite struct {
	suite.Suite
}

func TestFindPatientSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindPatientSuite))
}

func (s *FindPatientSuite) TestHandle() {
	type Sut struct {
		Sut                controller.FindPatient
		Input              query.FindPatientInput
		FindPatientHandler *mocks.Handler[*query.FindPatientInput, *query.FindPatientOutput]
	}

	makeSut := func() *Sut {
		input := fake.FindPatientInput()
		findPatientHandler := mocks.NewHandler[*query.FindPatientInput, *query.FindPatientOutput](
			s.T(),
		)
		sut := controller.NewFindPatient(findPatientHandler)
		return &Sut{Sut: sut, FindPatientHandler: findPatientHandler, Input: *input}
	}

	s.Run("should find one patient successfuly", func() {
		sut := makeSut()

		sut.FindPatientHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&query.AllPatientsOutput{ID: sut.Input.ID}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.FindPatientHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
	})

	s.Run("invalid Preload", func() {
		sut := makeSut()

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Queries: gintest.BuildScopeQuery(
				gintest.PreloadOption(sut.Input.Preload.Add("invalid").Slice()),
			),
			Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("ID", func() {
		s.Run("should return error when empty", func() {
			sut := makeSut()

			sut.Input.ID = 0

			ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
			})

			s.Equal(http.StatusBadRequest, ctx.Writer.Status())
		})
	})

	s.Run("panics when FindPatientHandler.Handle returns error", func() {
		sut := makeSut()

		sut.FindPatientHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
			})
		})
	})
}
