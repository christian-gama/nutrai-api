package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// DurationInWeeks is a data type that specifies the length of a specific diet plan, measured
// in weeks.
type DurationInWeeks int16

// Int16 returns the int16 representation of the duration.
func (d DurationInWeeks) Int16() int16 {
	return int16(d)
}

// Validate returns an error if the duration is invalid.
func (d DurationInWeeks) Validate() error {
	const fieldName = "DurationInWeeks"
	const maxWeeks = 520
	const minWeeks = 1

	if d < minWeeks {
		return errutil.Invalid(
			fieldName,
			fmt.Sprintf("cannot be less than %d week", minWeeks),
		)
	}

	if d > 520 {
		return errutil.Invalid(
			fieldName,
			fmt.Sprintf("cannot be greater than %d weeks", maxWeeks),
		)
	}

	return nil
}
