package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// WeightKG is a data type that denotes a patient's weight in kilograms. It is a vital measure in
// health assessments and planning diet routines.
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
		return errutil.Required(fieldName)
	}

	if w <= minWeight {
		return errutil.Invalid(fieldName, "cannot be less than %d", minWeight)
	}

	if w >= maxWeight {
		return errutil.Invalid(fieldName, "cannot be greater than %d", maxWeight)
	}

	return nil
}
