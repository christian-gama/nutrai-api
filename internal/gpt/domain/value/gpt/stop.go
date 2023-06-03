package value

// Stop is a data type that represents the stop condition.
type Stop string

// String returns the string representation of the Stop.
func (n Stop) String() string {
	return string(n)
}
