package command

import value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"

// CheckCredentialsInput is the input for CheckCredentials.
type CheckCredentialsInput struct {
	Email    value.Email    `json:"email" faker:"email" validate:"required,email"`
	Password value.Password `json:"password" faker:"len=8" validate:"required,user_password"`
}
