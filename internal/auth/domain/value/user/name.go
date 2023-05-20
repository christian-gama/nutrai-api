package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Name is a simple data type encapsulating a user's name. It serves as a personal
// identifier within a system.
type Name string

// String returns the string representation of the name.
func (n Name) String() string {
	return string(n)
}

// Validate returns an error if the name is invalid.
func (n Name) Validate() error {
	const fieldName = "Name"
	const minLen = 2
	const maxLen = 150

	if n == "" {
		return errutil.NewErrRequired(fieldName)
	}

	if len(n.String()) < minLen {
		return errutil.NewErrInvalid(
			fieldName,
			fmt.Sprintf("cannot be less than %d characters", minLen),
		)
	}

	if len(n.String()) > maxLen {
		return errutil.NewErrInvalid(
			fieldName,
			fmt.Sprintf("cannot be greater than %d characters", maxLen),
		)
	}

	return nil
}
