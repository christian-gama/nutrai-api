package command

import (
	sharedvalue "github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// ChangePasswordInput represents the input data for the ChangePassword command.
type ChangePasswordInput struct {
	ID sharedvalue.ID `form:"id" faker:"uint"`

	Password value.Password `json:"password" faker:"len=8"`
}
