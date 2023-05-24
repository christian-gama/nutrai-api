package value

// Body represents the body of an email message.
type Body string

// String returns the string representation of the body.
func (b Body) String() string {
	return string(b)
}
