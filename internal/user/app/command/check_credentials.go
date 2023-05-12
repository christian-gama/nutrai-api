package command

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// CheckCredentialsInput is the input for CheckCredentials.
type CheckCredentialsInput struct {
	Email    value.Email    `json:"email"`
	Password value.Password `json:"password"`
}
