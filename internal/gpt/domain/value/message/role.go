package value

import "github.com/christian-gama/nutrai-api/pkg/errutil/errors"

// Role is a data type that represents the role of a message.
type Role string

const (
	User      Role = "user"
	System    Role = "system"
	Assistant Role = "assistant"
)

// String returns the string representation of the Role.
func (n Role) String() string {
	return string(n)
}

// Validate returns an error if the role is invalid.
func (n Role) Validate() error {
	const fieldName = "Role"

	if len(n) == 0 {
		return errors.Required(fieldName)
	}

	return nil
}
