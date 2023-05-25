package value

// Email is a simple data type that represents an email address.
// It serves as a key contact detail for identifying an entity.
type Email string

// String returns the string representation of the email.
func (e Email) String() string {
	return string(e)
}
