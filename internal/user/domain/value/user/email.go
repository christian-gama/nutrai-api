package value

import (
	"net/mail"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Email represents a user email.
type Email string

// String returns the string representation of the email.
func (e Email) String() string {
	return string(e)
}

// Validate returns an error if the email is invalid.
func (e Email) Validate() error {
	const fieldName = "email"

	if e == "" {
		return errutil.NewErrRequired(fieldName)
	}

	if _, err := mail.ParseAddress(e.String()); err != nil {
		return errutil.NewErrInvalid(fieldName, err.Error())
	}

	return nil
}
