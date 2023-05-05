package value

// Name is the name of a diet. It is used to identify a diet.
type Name string

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// IsValid returns true if the Name is valid.
func (n Name) IsValid() bool {
	return len(n) > 0 && len(n) < 100
}
