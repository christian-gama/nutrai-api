package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Stop is a data type that represents the stop condition.
type Stop string

// String returns the string representation of the Stop.
func (n Stop) String() string {
	return string(n)
}

// Validate returns an error if the stop is invalid.
func (n Stop) Validate() error {
	const fieldName = "Stop"

	if len(n) == 0 {
		return errors.Required(fieldName)
	}

	return nil
}
