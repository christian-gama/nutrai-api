package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// TopP is a data type that represents the top p of a model.
type TopP float32

// Float32 returns the string representation of the TopP.
func (n TopP) Float32() float32 {
	return float32(n)
}

// Validate returns an error if the top p is invalid.
func (n TopP) Validate() error {
	const fieldName = "TopP"

	if n < 0 || n > 1 {
		return errors.Invalid(
			fieldName,
			"must be greater than or equal to 0 and less than or equal to 1",
		)
	}

	return nil
}
