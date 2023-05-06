package value

import (
	"fmt"

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
	const fieldName = "password"
	const minLen = 8
	const maxLen = 32

	if p == "" {
		return errutil.NewErrRequired(fieldName)
	}

	if len(p.String()) < minLen {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be less than %d characters", minLen))
	}

	if len(p.String()) > maxLen {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be greater than %d characters", maxLen))
	}

	return nil
}
