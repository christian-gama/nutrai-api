package value

import (
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// ID is a data type that serves as a unique identifier for a specific domain entity, such as a user.
// It is crucial for managing and referencing entities within the system.
type ID uint

// Uint returns the ID as an unsigned integer.
func (i ID) Uint() uint {
	return uint(i)
}

// Validate returns true if the ID is valid.
func (i ID) Validate() error {
	if i == 0 {
		return errutil.NewErrRequired("id")
	}

	return nil
}
