package command

import (
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

// SaveUserInput represents the input data for the SaveUser command.
type SaveUserInput struct {
	Email    userValue.Email    `json:"email" validate:"required,email" faker:"email"`
	Password userValue.Password `json:"password" validate:"required,alphanum,min=8,max=32" faker:"len=8"`
	Name     userValue.Name     `json:"name" validate:"required,min=2,max=100" faker:"name"`
}
