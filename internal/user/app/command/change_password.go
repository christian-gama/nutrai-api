package command

import (
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// ChangePasswordInput represents the input data for the ChangePassword command.
type ChangePasswordInput struct {
	ID coreValue.ID `uri:"id" faker:"uint"`

	Password value.Password `json:"password" faker:"len=8"`
}
