package value

import (
	"net/mail"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Email is a simple data type that represents a user's email address.
// It serves as a key contact detail for a user in a system.
type Email string

// String returns the string representation of the email.
func (e Email) String() string {
	return string(e)
}

// Validate returns an error if the email is invalid.
func (e Email) Validate() error {
	const fieldName = "Email"

	if e == "" {
		return errutil.NewErrRequired(fieldName)
	}

	if _, err := mail.ParseAddress(e.String()); err != nil {
		// err.Error output is equal to `mail: error message`
		msg := strings.Split(err.Error(), ":")[1]
		msg = strings.TrimSpace(msg)

		return errutil.NewErrInvalid(fieldName, msg)
	}

	return nil
}
