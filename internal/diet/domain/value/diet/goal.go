package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Goal represents the goal of the diet.
type Goal string

// String returns the string representation of the goal.
func (g Goal) String() string {
	return string(g)
}

// Validate returns an error if the goal is invalid.
func (g Goal) Validate() error {
	const fieldName = "goal"

	validGoals := []Goal{
		WeightLoss,
		WeightGain,
		Maintain,
		ImprovedHealth,
		IncreaseEnergy,
		MuscleGain,
	}

	if len(g) == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	for _, validGoal := range validGoals {
		if validGoal == g {
			return nil
		}
	}

	return errutil.NewErrInvalid(fieldName, fmt.Sprintf("must be one of %v", validGoals))
}

const (
	WeightLoss     Goal = "WEIGHT_LOSS"
	WeightGain     Goal = "WEIGHT_GAIN"
	Maintain       Goal = "MAINTAIN"
	ImprovedHealth Goal = "IMPROVED_HEALTH"
	IncreaseEnergy Goal = "INCREASE_ENERGY"
	MuscleGain     Goal = "MUSCLE_GAIN"
)