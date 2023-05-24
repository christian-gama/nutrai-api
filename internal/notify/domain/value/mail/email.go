package value

import value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"

// Email represents an email address.
type Email value.Email

// String returns the string representation of the email address.
func (e Email) String() string {
	return string(e)
}
