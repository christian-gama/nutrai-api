package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

type Allergy string

func (a Allergy) String() string {
	return string(a)
}

func (a Allergy) Validate() error {
	const fieldName = "allergy"
	const maxLen = 100

	if len(a) > maxLen {
		return errutil.NewErrInvalid(fieldName, "must be less than %d characters", maxLen)
	}

	return nil
}
