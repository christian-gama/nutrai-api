package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/hasher"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/app/command"
)

// ChangePasswordHandler represents the ChangePassword command.
type ChangePasswordHandler = command.Handler[*ChangePasswordInput]

// ChangePasswordHandlerImpl represents the implementation of the ChangePassword command.
type ChangePasswordHandlerImpl struct {
	repo.User
	hasher.Hasher
}

// NewChangePasswordHandler returns a new Save instance.
func NewChangePasswordHandler(userRepo repo.User, hasher hasher.Hasher) ChangePasswordHandler {
	if userRepo == nil {
		panic(errors.New("repo.User cannot be nil"))
	}

	if hasher == nil {
		panic(errors.New("hasher.Hasher cannot be nil"))
	}

	return &ChangePasswordHandlerImpl{userRepo, hasher}
}

// Handle implements command.Handler.
func (c *ChangePasswordHandlerImpl) Handle(ctx context.Context, input *ChangePasswordInput) error {
	savedUser, err := c.Find(ctx, repo.FindUserInput{ID: input.User.ID})
	if err != nil {
		return err
	}

	hashedPassword, err := c.Hash(input.Password)
	if err != nil {
		return err
	}

	user, err := savedUser.
		SetPassword(hashedPassword).
		Validate()
	if err != nil {
		return err
	}

	return c.Update(ctx, repo.UpdateUserInput{User: user, ID: user.ID})
}
