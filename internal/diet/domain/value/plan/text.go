package value

// Text is a data type that represents the description of a specific diet plan.
type Text string

// String returns the string representation of the Plan.
func (p Text) String() string {
	return string(p)
}
