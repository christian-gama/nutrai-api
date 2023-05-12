package command

import "github.com/christian-gama/nutrai-api/internal/shared/domain/value"

// DeleteUserInput represents the input data for the DeleteUser command.
type DeleteUserInput struct {
	ID value.ID `form:"id" faker:"uint"`
}
