package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Password is a data type that encapsulates a user's secure login credential within a system.
type Password string

// String returns the string representation of the password.
func (p Password) String() string {
	return string(p)
}

// Validate returns an error if the password is invalid.
func (p Password) Validate() error {
	const fieldName = "Password"

	if p == "" {
		return errutil.Required(fieldName)
	}

	return nil
}
