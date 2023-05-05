package value

// Description is a diet description.
type Description string

// String returns the string representation of the Description.
func (n Description) String() string {
	return string(n)
}

func (n Description) IsValid() bool {
	return len(n) > 0 && len(n) < 500
}
