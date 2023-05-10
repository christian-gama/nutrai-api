package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Password represents a user password.
type Password string

// String returns the string representation of the password.
func (p Password) String() string {
	return string(p)
}

// Validate returns an error if the password is invalid.
func (p Password) Validate() error {
	const fieldName = "Password"

	if p == "" {
		return errutil.NewErrRequired(fieldName)
	}

	return nil
}
