package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
)

// DeleteMeInput represents the input data for the DeleteMe command.
type DeleteMeInput struct {
	User *user.User `ctx:"current_user" json:"-"`
}
