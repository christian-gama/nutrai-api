package value

// Content is a data type that represents the content of a message.
type Content string

// String returns the string representation of the Content.
func (n Content) String() string {
	return string(n)
}
