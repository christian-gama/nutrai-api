package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindPlanHandlerSuite struct {
	suite.Suite
}

func TestFindPlanHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindPlanHandlerSuite))
}

func (s *FindPlanHandlerSuite) TestPlanHandler() {
	type Mock struct {
		PlanRepo *mocks.Plan
	}

	type Sut struct {
		Sut   query.FindPlanHandler
		Ctx   context.Context
		Input *query.FindPlanInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PlanRepo: mocks.NewPlan(s.T()),
		}

		input := queryFake.FindPlanInput()

		sut := query.NewFindPlanHandler(mock.PlanRepo)

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

			sut.Mock.PlanRepo.On("Find", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(output)
			s.ErrorIs(assert.AnError, err)
		})
	})

	s.Run("PlanHandler (Success)", func() {
		s.Run("Should return a plan", func() {
			sut := makeSut()

			plan := fake.Plan()
			plan.ID = sut.Input.ID

			sut.Mock.PlanRepo.On("Find", sut.Ctx, mock.Anything).
				Return(plan, nil)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Require().NotNil(output)
			s.Equal(sut.Input.ID, output.ID, "ID should be the same")
		})
	})
}
