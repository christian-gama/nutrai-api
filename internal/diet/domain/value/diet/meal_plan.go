package value

// MealPlan represents the meal plan of a diet.
type MealPlan string

// String returns the string representation of the meal plan.
func (m MealPlan) String() string {
	return string(m)
}

// IsValid returns true if the meal plan is valid.
func (m MealPlan) IsValid() bool {
	switch m {
	case Ketogenic, Vegetarian, Vegan, Mediterranean, Paleolithic, LowCarb:
		return true
	}
	return false
}

const (
	Ketogenic     MealPlan = "ketogenic"
	Vegetarian    MealPlan = "vegetarian"
	Vegan         MealPlan = "vegan"
	Mediterranean MealPlan = "mediterranean"
	Paleolithic   MealPlan = "paleolithic"
	LowCarb       MealPlan = "low-carb"
)
