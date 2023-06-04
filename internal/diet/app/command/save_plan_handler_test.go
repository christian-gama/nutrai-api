package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	cmdFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/command"
	dietFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	planFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	dietRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SavePlanHandlerSuite struct {
	suite.Suite
}

func TestSavePlanHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SavePlanHandlerSuite))
}

func (s *SavePlanHandlerSuite) TestSaveHandler() {
	type Mock struct {
		PlanRepo *dietRepoMock.Plan
		DietRepo *dietRepoMock.Diet
	}

	type Sut struct {
		Sut   command.SavePlanHandler
		Ctx   context.Context
		Input *command.SavePlanInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PlanRepo: dietRepoMock.NewPlan(s.T()),
			DietRepo: dietRepoMock.NewDiet(s.T()),
		}

		input := cmdFake.SavePlanInput()

		sut := command.NewSavePlanHandler(mock.PlanRepo, mock.DietRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("TestSaveHandler (Error)", func() {
		s.Run("Should return error when saving plan fails", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Find", sut.Ctx, mock.Anything).
				Return(dietFake.Diet(), nil)

			sut.Mock.PlanRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
		})

		s.Run("Should return error when validating plan fails", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Find", sut.Ctx, mock.Anything).
				Return(dietFake.Diet(), nil)

			sut.Mock.PlanRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
		})
	})

	s.Run("TestSaveHandler (Success)", func() {
		s.Run("Should return nil when saving plan succeeds", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Find", sut.Ctx, mock.Anything).
				Return(dietFake.Diet(), nil)

			sut.Mock.PlanRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(planFake.Plan(), nil)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
		})
	})
}
