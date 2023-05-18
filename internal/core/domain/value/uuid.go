package value

// UUID is a simple data type that represents a universally unique identifier.
type UUID string

// String returns the string representation of the UUID.
func (u UUID) String() string {
	return string(u)
}
