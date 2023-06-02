package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Content is a data type that represents the content of a message.
type Content string

// String returns the string representation of the Content.
func (n Content) String() string {
	return string(n)
}

// Validate returns an error if the content is invalid.
func (n Content) Validate() error {
	const fieldName = "Content"

	if len(n) == 0 {
		return errors.Required(fieldName)
	}

	return nil
}
