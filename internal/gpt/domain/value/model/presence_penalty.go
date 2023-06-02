package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// PresencePenalty is a data type that represents the presence penalty.
type PresencePenalty float32

// Float32 returns the string representation of the PresencePenalty.
func (n PresencePenalty) Float32() float32 {
	return float32(n)
}

// Validate returns an error if the presence penalty is invalid.
func (n PresencePenalty) Validate() error {
	const fieldName = "PresencePenalty"

	if n < -2 || n > 2 {
		return errors.Invalid(
			fieldName,
			"must be greater than or equal to -2 and less than or equal to 2",
		)
	}

	return nil
}
