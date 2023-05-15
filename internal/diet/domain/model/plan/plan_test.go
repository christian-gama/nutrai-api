package plan_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/plan"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/diet/domain/model/plan"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type PlanSuite struct {
	suite.Suite
}

func TestPlanSuite(t *testing.T) {
	suite.RunUnitTest(t, new(PlanSuite))
}

func (s *PlanSuite) TestNewPlan() {
	type Sut struct {
		Sut  func() (*plan.Plan, error)
		Data *plan.Plan
	}

	makeSut := func() *Sut {
		data := fake.Plan()

		sut := func() (*plan.Plan, error) {
			return plan.NewBuilder().
				SetID(data.ID).
				SetText(data.Text).
				Build()
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewPlan (Error)", func() {
		s.Run("ID", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.ID = 0

				plan, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(plan)
			})
		})

		s.Run("Text", func() {
			s.Run("Should return an error when empty", func() {
				sut := makeSut()
				sut.Data.Text = ""

				plan, err := sut.Sut()

				s.ErrorAsRequired(err)
				s.Nil(plan)
			})
		})
	})

	s.Run("TestNewPlan (Success)", func() {
		s.Run("Should return a plan when valid", func() {
			sut := makeSut()

			plan, err := sut.Sut()

			s.NoError(err)
			s.NotNil(plan)
			s.Equal(sut.Data.ID, plan.ID, "should have the same id")
			s.Equal(sut.Data.Text, plan.Text, "should have the same text")
		})
	})
}
