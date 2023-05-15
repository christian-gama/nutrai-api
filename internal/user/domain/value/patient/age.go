package value

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Age is a data type that signifies a patient's age. It's crucial for determining
// appropriate health and diet plans.
type Age int8

// Int8 returns the int8 representation of the age.
func (w Age) Int8() int8 {
	return int8(w)
}

// Validate returns an error if the age is invalid.
func (w Age) Validate() error {
	const fieldName = "Age"
	const minAge = 1
	const maxAge = 120

	if w == 0 {
		return errutil.NewErrRequired(fieldName)
	}

	if w < minAge {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be less than %d", minAge))
	}

	if w >= maxAge {
		return errutil.NewErrInvalid(fieldName, fmt.Sprintf("cannot be greater than %d", maxAge))
	}

	return nil
}
