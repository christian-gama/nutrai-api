package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
)

// DeletePlanInput represents the input data for the DeletePlan command.
type DeletePlanInput struct {
	User *user.User `ctx:"current_user" json:"-"`

	ID uint `uri:"id" validate:"required" faker:"uint"`
}
