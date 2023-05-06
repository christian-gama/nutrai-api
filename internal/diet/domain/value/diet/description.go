package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Description is a diet description.
type Description string

// String returns the string representation of the Description.
func (n Description) String() string {
	return string(n)
}

// Validate returns an error if the description is invalid.
func (n Description) Validate() error {
	const fieldName = "description"
	const maxChars = 500

	if len(n) == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if len(n) > maxChars {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be longer than %d characters", maxChars))
	}

	return nil
}
