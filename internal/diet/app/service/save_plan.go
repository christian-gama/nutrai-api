package service

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	value "github.com/christian-gama/nutrai-api/internal/diet/domain/value/plan"
)

type SavePlanInput struct {
	User *user.User `ctx:"current_user" json:"-"`

	DietID coreValue.ID `json:"dietID" validate:"required" faker:"uint"`
}

type SavePlanOutput struct {
	ID     coreValue.ID `json:"id" faker:"uint"`
	DietID coreValue.ID `json:"dietID" faker:"uint"`
	Text   value.Text   `json:"text" faker:"sentence"`
}
