package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/model/diet"
	"github.com/christian-gama/nutrai-api/internal/diet/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/diet"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// SaveDietHandler represents the SaveDiet command.
type SaveDietHandler = command.Handler[*SaveDietInput]

// saveDietHandlerImpl represents the implementation of the SaveDiet command.
type saveDietHandlerImpl struct {
	repo.Diet
}

// NewSaveDietHandler returns a new Save instance.
func NewSaveDietHandler(dietRepo repo.Diet) SaveDietHandler {
	errutil.MustBeNotEmpty("repo.Diet", dietRepo)

	return &saveDietHandlerImpl{dietRepo}
}

// Handle implements command.Handler.
func (c *saveDietHandlerImpl) Handle(ctx context.Context, input *SaveDietInput) error {
	data, err := diet.NewDiet().
		SetName(input.Name).
		SetDescription(input.Description).
		SetDurationInWeeks(input.DurationInWeeks).
		SetGoal(input.Goal).
		SetMealPlan(input.MealPlan).
		SetMonthlyCostUSD(input.MonthlyCostUSD).
		SetPatientID(input.PatientID).
		SetRestrictedFood(
			slice.
				Map(
					input.RestrictedFood,
					func(restrictedFood value.RestrictedFood) *diet.RestrictedFood {
						return diet.NewRestrictedFood().SetName(restrictedFood)
					},
				).
				Build(),
		).
		Validate()
	if err != nil {
		return err
	}

	return command.Response(c.Save(ctx, repo.SaveDietInput{Diet: data}))
}
