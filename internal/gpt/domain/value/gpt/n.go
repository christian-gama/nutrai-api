package value

// N is a data type that represents the count of responses.
type N int

// Int returns the string representation of the N.
func (n N) Int() int {
	return int(n)
}
