package command

import (
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// ChangePasswordInput represents the input data for the ChangePassword command.
type ChangePasswordInput struct {
	ID coreValue.ID `uri:"id" faker:"uint"`

	Password value.Password `json:"password" faker:"len=8"`
}
