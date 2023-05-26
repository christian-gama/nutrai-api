package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Plan is a data type that represents the description of a specific diet plan.
type Plan string

// String returns the string representation of the Plan.
func (p Plan) String() string {
	return string(p)
}

// Validate returns an error if the plan is invalid.
func (p Plan) Validate() error {
	const fieldName = "Plan"
	const maxChars = 1000

	if len(p) == 0 {
		return errors.Required(fieldName)
	}

	if len(p) > maxChars {
		return errors.Invalid(
			fieldName,
			fmt.Sprintf("cannot be longer than %d characters", maxChars),
		)
	}

	return nil
}
