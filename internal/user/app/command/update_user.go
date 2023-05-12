package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// UpdateUserInput represents the input data for the UpdateUser command.
type UpdateUserInput struct {
	Email value.Email `json:"email" faker:"email"`
	Name  value.Name  `json:"name" faker:"name"`
}
