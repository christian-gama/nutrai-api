package value

type MaxTokens int

// Int returns the string representation of the MaxTokens.
func (n MaxTokens) Int() int {
	return int(n)
}
