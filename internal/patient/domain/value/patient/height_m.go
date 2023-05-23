package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// HeightM is a data type that represents a patient's height in meters. It is an essential
// factor in calculating BMI and determining health and diet plans.
type HeightM float32

// Float32 returns the float32 representation of the height.
func (w HeightM) Float32() float32 {
	return float32(w)
}

// Validate returns an error if the height is invalid.
func (w HeightM) Validate() error {
	const fieldName = "HeightM"
	const minHeight = 1
	const maxHeight = 3

	if w == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if w <= minHeight {
		return errutil.NewErrInvalid(fieldName, "cannot be less than %d", minHeight)
	}

	if w >= maxHeight {
		return errutil.NewErrInvalid(fieldName, "cannot be greater than %d", maxHeight)
	}

	return nil
}
