package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// DeleteUserHandler represents the DeleteUser command.
type DeleteUserHandler = command.Handler[*DeleteUserInput]

// deleteUserHandlerImpl represents the implementation of the DeleteUser command.
type deleteUserHandlerImpl struct {
	repo.User
}

// NewDeleteUserHandler returns a new Delete instance.
func NewDeleteUserHandler(r repo.User) DeleteUserHandler {
	if r == nil {
		panic(errors.New("repo.User cannot be nil"))
	}

	return &deleteUserHandlerImpl{r}
}

// Handle implements command.Handler.
func (c *deleteUserHandlerImpl) Handle(ctx context.Context, input *DeleteUserInput) error {
	_, err := c.Find(ctx, repo.FindUserInput{ID: input.ID})
	if err != nil {
		return err
	}

	return c.Delete(ctx, repo.DeleteUserInput{IDs: []value.ID{input.ID}})
}
