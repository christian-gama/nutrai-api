package command

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// SavePlanInput represents the input data for the SavePlan command.
type SavePlanInput struct {
	User *user.User `ctx:"current_user" json:"-"`

	DietID coreValue.ID `json:"dietID" validate:"required" faker:"uint"`
}
