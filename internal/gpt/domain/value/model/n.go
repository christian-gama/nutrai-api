package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// N is a data type that represents the count of responses.
type N int

// Int returns the string representation of the N.
func (n N) Int() int {
	return int(n)
}

// Validate returns an error if the n is invalid.
func (n N) Validate() error {
	const fieldName = "N"

	if n <= 0 {
		return errors.Invalid(fieldName, "must be greater than to 0")
	}

	return nil
}
