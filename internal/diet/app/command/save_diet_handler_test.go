package command_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/command"
	fakeDiet "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	userRepoMock "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SaveDietHandlerSuite struct {
	suite.Suite
}

func TestSaveDietHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(SaveDietHandlerSuite))
}

func (s *SaveDietHandlerSuite) TestSaveHandler() {
	type Mock struct {
		DietRepo *userRepoMock.Diet
	}

	type Sut struct {
		Sut   command.SaveDietHandler
		Ctx   context.Context
		Input *command.SaveDietInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			DietRepo: userRepoMock.NewDiet(s.T()),
		}

		input := fake.SaveDietInput()

		sut := command.NewSaveDietHandler(mock.DietRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("TestSaveHandler (Error)", func() {
		s.Run("Should return error when saving diet fails", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
		})

		s.Run("Should return error when validating diet fails", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.ErrorIs(err, assert.AnError)
		})
	})

	s.Run("TestSaveHandler (Success)", func() {
		s.Run("Should return nil when saving diet succeeds", func() {
			sut := makeSut()

			sut.Mock.DietRepo.
				On("Save", sut.Ctx, mock.Anything).
				Return(fakeDiet.Diet(), nil)

			err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
		})
	})
}
