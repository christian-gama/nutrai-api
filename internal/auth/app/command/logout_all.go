package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
)

// LogoutInput represents the input data for the LogoutInput command.
type LogoutAllInput struct {
	User *user.User `ctx:"currentUser" json:"-"`
}
