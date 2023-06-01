package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Description is a data type that represents a detailed explanation or outline of
// a specific diet plan.
type Description string

// String returns the string representation of the Description.
func (n Description) String() string {
	return string(n)
}

// Validate returns an error if the description is invalid.
func (n Description) Validate() error {
	const fieldName = "Description"
	const maxChars = 500

	if len(n) == 0 {
		return errors.Required(fieldName)
	}

	if len(n) > maxChars {
		return errors.Invalid(
			fieldName,
			fmt.Sprintf("cannot be longer than %d characters", maxChars),
		)
	}

	return nil
}
