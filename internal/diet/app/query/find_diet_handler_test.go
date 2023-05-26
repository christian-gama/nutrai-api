package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FindDietHandlerSuite struct {
	suite.Suite
}

func TestFindDietHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(FindDietHandlerSuite))
}

func (s *FindDietHandlerSuite) TestDietHandler() {
	type Mock struct {
		DietRepo *mocks.Diet
	}

	type Sut struct {
		Sut   query.FindDietHandler
		Ctx   context.Context
		Input *query.FindDietInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			DietRepo: mocks.NewDiet(s.T()),
		}

		input := queryFake.FindDietInput()

		sut := query.NewFindDietHandler(mock.DietRepo)

		return Sut{
			Sut:   sut,
			Ctx:   context.Background(),
			Input: input,
			Mock:  mock,
		}
	}

	s.Run("DietHandler (Error)", func() {
		s.Run("Should return an error if the repo returns an error", func() {
			sut := makeSut()

			sut.Mock.DietRepo.On("Find", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(output)
			s.ErrorIs(assert.AnError, err)
		})
	})

	s.Run("DietHandler (Success)", func() {
		s.Run("Should return a diet", func() {
			sut := makeSut()

			diet := fake.Diet()
			diet.ID = sut.Input.ID

			sut.Mock.DietRepo.On("Find", sut.Ctx, mock.Anything).
				Return(diet, nil)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Require().NotNil(output)
			s.Equal(sut.Input.ID, output.ID, "ID should be the same")
		})
	})
}
