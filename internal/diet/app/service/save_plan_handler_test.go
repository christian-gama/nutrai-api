package service_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/service"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	svcFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/service"
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
		Sut   service.SavePlanHandler
		Ctx   context.Context
		Input *service.SavePlanInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			PlanRepo: dietRepoMock.NewPlan(s.T()),
			DietRepo: dietRepoMock.NewDiet(s.T()),
		}

		input := svcFake.SavePlanInput()

		sut := service.NewSavePlanHandler(mock.PlanRepo, mock.DietRepo)

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

			result, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
			s.Nil(result)
		})

		s.Run("Should return error when validating plan fails", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Find", sut.Ctx, mock.Anything).
				Return(dietFake.Diet(), nil)

			sut.Mock.PlanRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			result, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
			s.Nil(result)
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

			result, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotNil(result)
		})

		s.Run("Should call Save with correct Diet data", func() {
			sut := makeSut()

			diet := dietFake.Diet()
			diet.ID = sut.Input.DietID
			sut.Mock.DietRepo.
				On("Find", sut.Ctx, mock.Anything).
				Return(diet, nil)

			sut.Mock.PlanRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(planFake.Plan(), nil)

			result, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.NotNil(result)
			sut.Mock.PlanRepo.AssertCalled(
				s.T(),
				"Save",
				sut.Ctx,
				mock.MatchedBy(func(r repo.SavePlanInput) bool {
					return r.Plan.DietID == diet.ID
				}),
			)
		})
	})
}
