package value

// Password is a data type that encapsulates a user's secure login credential within a system.
type Password string

// String returns the string representation of the password.
func (p Password) String() string {
	return string(p)
}
