package value

// Role is a data type that represents the role of a message.
type Role string

// String returns the string representation of the Role.
func (n Role) String() string {
	return string(n)
}
