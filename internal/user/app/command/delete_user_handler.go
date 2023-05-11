package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// DeleteUserHandler represents the DeleteUser command.
type DeleteUserHandler = command.Handler[*DeleteUserInput]

// deleteUserHandlerImpl represents the implementation of the DeleteUser command.
type deleteUserHandlerImpl struct {
	repo.User
}

// NewDeleteUserHandler returns a new Delete instance.
func NewDeleteUserHandler(p repo.User) DeleteUserHandler {
	return &deleteUserHandlerImpl{p}
}

// Handle implements command.Handler.
func (c *deleteUserHandlerImpl) Handle(ctx context.Context, input *DeleteUserInput) error {
	_, err := c.Find(ctx, repo.FindUserInput{ID: value.ID(input.ID)})
	if err != nil {
		return err
	}

	return c.Delete(ctx, repo.DeleteUserInput{IDs: []value.ID{value.ID(input.ID)}})
}
