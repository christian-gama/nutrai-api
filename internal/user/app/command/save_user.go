package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// SaveUserInput represents the input data for the SaveUser command.
type SaveUserInput struct {
	Email    value.Email    `json:"email" faker:"email"`
	Password value.Password `json:"password" faker:"len=8"`
	Name     value.Name     `json:"name" faker:"name"`
}
