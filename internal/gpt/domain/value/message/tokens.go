package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Tokens is a data type that represents the count tokens of a message.
type Tokens int

// Int returns the string representation of the Tokens.
func (n Tokens) Int() int {
	return int(n)
}

// Validate returns an error if the tokens is invalid.
func (n Tokens) Validate() error {
	const fieldName = "Tokens"

	if n < 0 {
		return errors.Invalid(fieldName, "must be greater than or equal to 0")
	}

	return nil
}
