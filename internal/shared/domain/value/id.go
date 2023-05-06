package value

import "github.com/christian-gama/nutrai-api/pkg/errutil"

// ID represents the ID of a domain entity.
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
