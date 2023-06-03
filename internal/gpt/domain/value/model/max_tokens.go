package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

type MaxTokens int

// Int returns the string representation of the MaxTokens.
func (n MaxTokens) Int() int {
	return int(n)
}

// Validate returns an error if the max tokens is invalid.
func (n MaxTokens) Validate() error {
	const fieldName = "MaxTokens"

	if n <= 0 {
		return errors.Required(fieldName)
	}

	return nil
}
