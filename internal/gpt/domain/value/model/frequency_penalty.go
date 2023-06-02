package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// FrequencyPenalty is a data type that represents the frequency penalty.
type FrequencyPenalty float32

// Float32 returns the string representation of the FrequencyPenalty.
func (n FrequencyPenalty) Float32() float32 {
	return float32(n)
}

// Validate returns an error if the frequency penalty is invalid.
func (n FrequencyPenalty) Validate() error {
	const fieldName = "FrequencyPenalty"

	if n < -2 || n > 2 {
		return errors.Invalid(
			fieldName,
			"must be greater than or equal to -2 and less than or equal to 2",
		)
	}

	return nil
}
