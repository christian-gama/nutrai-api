package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
)

// LogoutInput represents the input data for the LogoutInput command.
type LogoutInput struct {
	User    *user.User  `ctx:"currentUser" json:"-"`
	Refresh value.Token `json:"refresh" validate:"required,jwt" faker:"jwt"`
}
