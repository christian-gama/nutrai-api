package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// WeightKG represents a patient weight in kilograms.
type WeightKG float32

// Float32 returns the float32 representation of the weight.
func (w WeightKG) Float32() float32 {
	return float32(w)
}

// Validate returns an error if the weight is invalid.
func (w WeightKG) Validate() error {
	const fieldName = "WeightKG"
	const minWeight = 1
	const maxWeight = 999

	if w == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if w <= minWeight {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be less than %d", minWeight))
	}

	if w >= maxWeight {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be greater than %d", maxWeight))
	}

	return nil
}
