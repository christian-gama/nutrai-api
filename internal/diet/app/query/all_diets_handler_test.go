package query_test

import (
	"context"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	"github.com/christian-gama/nutrai-api/internal/diet/app/query"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	queryFake "github.com/christian-gama/nutrai-api/testutils/fake/diet/app/query"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/diet"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/diet/domain/repo"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AllDietsHandlerSuite struct {
	suite.Suite
}

func TestAllDietsHandlerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(AllDietsHandlerSuite))
}

func (s *AllDietsHandlerSuite) TestDietHandler() {
	type Mock struct {
		DietRepo *mocks.Diet
	}

	type Sut struct {
		Sut   query.AllDietsHandler
		Ctx   context.Context
		Input *query.AllDietsInput
		Mock  *Mock
	}

	makeSut := func() Sut {
		mock := &Mock{
			DietRepo: mocks.NewDiet(s.T()),
		}

		input := queryFake.AllDietsInput()

		sut := query.NewAllDietsHandler(mock.DietRepo)

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

			sut.Mock.DietRepo.On("All", sut.Ctx, mock.Anything).
				Return(nil, assert.AnError)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.Error(err)
			s.Nil(output)
			s.ErrorIs(err, assert.AnError)
		})
	})

	s.Run("DietHandler (Success)", func() {
		s.Run("Should return a DietOutput", func() {
			sut := makeSut()

			sut.Mock.DietRepo.On("All", sut.Ctx, mock.Anything).
				Return(&queryer.PaginationOutput[*diet.Diet]{
					Results: []*diet.Diet{fake.Diet()},
					Total:   1,
				}, nil)

			output, err := sut.Sut.Handle(sut.Ctx, sut.Input)

			s.NoError(err)
			s.Require().NotNil(output)
			s.Len(output.Results, 1)
			s.Equal(1, output.Total)
		})
	})
}
