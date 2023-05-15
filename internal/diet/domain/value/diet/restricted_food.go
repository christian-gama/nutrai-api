package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// RestrictedFood is a data type that signifies a certain food or food group that is to be
// avoided or limited in a specific diet plan.
type RestrictedFood string

// String returns the string representation of a RestrictedFood.
func (r RestrictedFood) String() string {
	return string(r)
}

// Validate returns an error if the restricted food is invalid.
func (r RestrictedFood) Validate() error {
	const fieldName = "RestrictedFood"
	const maxChars = 100

	if len(r) == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if len(r) > maxChars {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be longer than %d characters", maxChars))
	}

	return nil
}
