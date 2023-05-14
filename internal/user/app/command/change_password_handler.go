package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
)

// ChangePasswordHandler represents the ChangePassword command.
type ChangePasswordHandler = command.Handler[*ChangePasswordInput]

// ChangePasswordHandlerImpl represents the implementation of the ChangePassword command.
type ChangePasswordHandlerImpl struct {
	repo.User
	service.HashPasswordHandler
}

// NewChangePasswordHandler returns a new Save instance.
func NewChangePasswordHandler(p repo.User, h service.HashPasswordHandler) ChangePasswordHandler {
	return &ChangePasswordHandlerImpl{p, h}
}

// Handle implements command.Handler.
func (c *ChangePasswordHandlerImpl) Handle(ctx context.Context, input *ChangePasswordInput) error {
	savedUser, err := c.Find(ctx, repo.FindUserInput{ID: input.ID})
	if err != nil {
		return err
	}

	hashPasswordOutput, err := c.HashPasswordHandler.Handle(
		ctx,
		&service.HashPasswordInput{Password: input.Password},
	)
	if err != nil {
		return err
	}

	user, err := user.NewBuilder().
		SetName(savedUser.Name).
		SetEmail(savedUser.Email).
		SetPassword(hashPasswordOutput.Password).
		Build()
	if err != nil {
		return err
	}

	return c.Update(ctx, repo.UpdateUserInput{User: user, ID: input.ID})
}
