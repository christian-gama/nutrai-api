package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// SaveUserInput represents the input data for the SaveUser command.
type SaveUserInput struct {
	Email    value.Email    `json:"email" validate:"required,email" faker:"email"`
	Password value.Password `json:"password" validate:"required,alphanum,min=8,max=32" faker:"len=8"`
	Name     value.Name     `json:"name" validate:"required,min=2,max=100" faker:"name"`
}
