package value

// RestrictedFood removes the food from the diet.
type RestrictedFood string

// String returns the string representation of a RestrictedFood.
func (r RestrictedFood) String() string {
	return string(r)
}

// IsValid returns true if the RestrictedFood is valid.
func (r RestrictedFood) IsValid() bool {
	return len(r) > 0 && len(r) < 100
}
