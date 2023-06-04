package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AllPlansHandlerSuite struct {
	suite.Suite
}

func TestAllPlansHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllPlansHandlerSuite))
}

func (s *AllPlansHandlerSuite) TestPlanHandler() {
	type Mock struct {
		PlanRepo *mocks.Plan
	}

	type Sut struct {
		Sut   query.AllPlansHandler
		Ctx   context.Context
		Input *query.AllPlansInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PlanRepo: mocks.NewPlan(s.T()),
		}

		input := queryFake.AllPlansInput()

		sut := query.NewAllPlansHandler(mock.PlanRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("PlanHandler (Error)", func() {
		s.Run("Should return an error if the repo returns an error", func() {
			sut := makeSut()

			sut.Mock.PlanRepo.On("All", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(output)
			s.ErrorIs(err, assert.AnError)
		})
	})

	s.Run("PlanHandler (Success)", func() {
		s.Run("Should return a PlanOutput", func() {
			sut := makeSut()

			sut.Mock.PlanRepo.On("All", sut.Ctx, mock.Anything).
				Return(queryFake.AllPlansOutput(), nil)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Require().NotNil(output)
			s.Len(output.Results, 1)
			s.Equal(1, output.Total)
		})
	})
}
