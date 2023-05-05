package value

// Goal represents the goal of the diet.
type Goal string

// String returns the string representation of the goal.
func (g Goal) String() string {
	return string(g)
}

// IsValid returns true if the goal is valid.
func (g Goal) IsValid() bool {
	switch g {
	case WeightLoss, WeightGain, Maintain, ImprovedHealth, IncreaseEnergy, MuscleGain:
		return true
	}
	return false
}

const (
	WeightLoss     Goal = "WEIGHT_LOSS"
	WeightGain     Goal = "WEIGHT_GAIN"
	Maintain       Goal = "MAINTAIN"
	ImprovedHealth Goal = "IMPROVED_HEALTH"
	IncreaseEnergy Goal = "INCREASE_ENERGY"
	MuscleGain     Goal = "MUSCLE_GAIN"
)
