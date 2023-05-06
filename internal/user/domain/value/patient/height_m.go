package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// HeightM represents a patient height in kilograms.
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
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be less than %d", minHeight))
	}

	if w >= maxHeight {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be greater than %d", maxHeight))
	}

	return nil
}
