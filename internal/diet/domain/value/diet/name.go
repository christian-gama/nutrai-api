package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Name is the name of a diet. It is used to identify a diet.
type Name string

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// Validate returns an error if the name is invalid.
func (n Name) Validate() error {
	const fieldName = "Name"
	const maxChars = 100

	if len(n) == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if len(n) > maxChars {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be longer than %d characters", maxChars))
	}

	return nil
}
