package value

type Token string

// String returns the string representation of the token.
func (t Token) String() string {
	return string(t)
}
