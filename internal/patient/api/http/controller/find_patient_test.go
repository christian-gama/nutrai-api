package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/patient/api/http/controller"
	"github.com/christian-gama/nutrai-api/internal/patient/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/patient/app/query"
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
	type Mock struct {
		FindPatientHandler *mocks.Handler[*query.FindPatientInput, *query.FindPatientOutput]
	}

	type Sut struct {
		Sut   controller.FindPatient
		Input query.FindPatientInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		input := fake.FindPatientInput()
		mock := &Mock{
			FindPatientHandler: mocks.NewHandler[*query.FindPatientInput, *query.FindPatientOutput](
				s.T(),
			),
		}
		sut := controller.NewFindPatient(mock.FindPatientHandler)
		return &Sut{Sut: sut, Mock: mock, Input: *input}
	}

	s.Run("should find one patient successfuly", func() {
		sut := makeSut()

		sut.Mock.FindPatientHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(&query.AllPatientsOutput{ID: sut.Input.ID}, nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
		sut.Mock.FindPatientHandler.AssertCalled(s.T(), "Handle", mock.Anything, mock.Anything)
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

		sut.Mock.FindPatientHandler.
			On("Handle", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprintf("%v", sut.Input.ID)},
			})
		})
	})
}
