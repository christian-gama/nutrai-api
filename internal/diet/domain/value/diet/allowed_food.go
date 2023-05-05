package value

// AllowedFood represents a diet's allowed food.
type AllowedFood string

// String returns the string representation of the allowed food.
func (a AllowedFood) String() string {
	return string(a)
}

// IsValid returns true if the allowed food is valid.
func (a AllowedFood) IsValid() bool {
	return len(a) > 0 && len(a) < 100
}
