package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Temperature is a data type that represents the temperature of a model.
type Temperature float32

// Float32 returns the string representation of the Temperature.
func (n Temperature) Float32() float32 {
	return float32(n)
}

// Validate returns an error if the temperature is invalid.
func (n Temperature) Validate() error {
	const fieldName = "Temperature"

	if n < 0 || n > 1 {
		return errors.Invalid(
			fieldName,
			"must be greater than or equal to 0 and less than or equal to 1",
		)
	}

	return nil
}
