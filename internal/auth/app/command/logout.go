package command

import "github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"

// LogoutInput represents the input data for the LogoutInput command.
type LogoutInput struct {
	User *user.User `ctx:"currentUser" json:"-"`
}