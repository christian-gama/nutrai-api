package value

// Name is a simple data type encapsulating a user's name. It serves as a personal
// identifier within a system.
type Name string

// String returns the string representation of the name.
func (n Name) String() string {
	return string(n)
}
