package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/command"
	planFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	planMock "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
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

func (s *DeletePlanSuite) TestDeletePlan() {
	type Mock struct {
		PlanRepo *planMock.Plan
	}

	type Sut struct {
		Sut   command.DeletePlanHandler
		Ctx   context.Context
		Input *command.DeletePlanInput
		Mock  *Mock
	}

	makeSut := func() *Sut {
		mock := &Mock{
			PlanRepo: planMock.NewPlan(s.T()),
		}

		input := fake.DeletePlanInput()

		sut := command.NewDeletePlanHandler(mock.PlanRepo)

		return &Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("Should delete the plan if succeed", func() {
		sut := makeSut()

		plan := planFake.Plan()
		sut.Mock.PlanRepo.On("Find", sut.Ctx, mock.Anything).Return(plan, nil)
		sut.Mock.PlanRepo.On("Delete", sut.Ctx, mock.Anything).Return(nil)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.Nil(err)
	})

	s.Run("Should return an error if the plan does not exist", func() {
		sut := makeSut()

		sut.Mock.PlanRepo.On("Find", sut.Ctx, mock.Anything).Return(nil, assert.AnError)

		err := sut.Sut.Handle(sut.Ctx, sut.Input)

		s.ErrorIs(err, assert.AnError)
	})
}
