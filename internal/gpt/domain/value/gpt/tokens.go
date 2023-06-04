package value

// Tokens is a data type that represents the count tokens of a message.
type Tokens int

// Int returns the string representation of the Tokens.
func (n Tokens) Int() int {
	return int(n)
}
