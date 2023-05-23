package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// DeleteMeHandler represents the DeleteMe command.
type DeleteMeHandler = command.Handler[*DeleteMeInput]

// deleteMeHandlerImpl represents the implementation of the DeleteUser command.
type deleteMeHandlerImpl struct {
	repo.User
}

// NewDeleteMeHandler returns a new Delete instance.
func NewDeleteMeHandler(userRepo repo.User) DeleteMeHandler {
	if userRepo == nil {
		panic(errors.New("repo.User cannot be nil"))
	}

	return &deleteMeHandlerImpl{userRepo}
}

// Handle implements command.Handler.
func (c *deleteMeHandlerImpl) Handle(ctx context.Context, input *DeleteMeInput) error {
	_, err := c.Find(ctx, repo.FindUserInput{ID: input.User.ID})
	if err != nil {
		return err
	}

	return c.Delete(ctx, repo.DeleteUserInput{IDs: []value.ID{input.User.ID}})
}
