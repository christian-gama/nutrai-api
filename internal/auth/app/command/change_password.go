package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

// ChangePasswordInput represents the input data for the ChangePassword command.
type ChangePasswordInput struct {
	User *user.User `ctx:"currentUser" json:"-"`

	Password value.Password `json:"password" faker:"len=8" validate:"required,min=8,max=32"`
}
