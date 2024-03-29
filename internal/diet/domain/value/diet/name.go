package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Name is a data type that represents the title or identifier of a specific diet plan. It provides
// an easy reference for diet categorization and selection.
type Name string

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// Validate returns an error if the name is invalid.
func (n Name) Validate() error {
	const fieldName = "Name"

	if len(n) == 0 {
		return errors.Required(fieldName)
	}

	return nil
}
