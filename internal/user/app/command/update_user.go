package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// UpdateUserInput represents the input data for the UpdateUser command.
type UpdateUserInput struct {
	Email value.Email `json:"email" validate:"required,email" faker:"email"`
	Name  value.Name  `json:"name" validate:"required,min=2,max=100" faker:"name"`
}
