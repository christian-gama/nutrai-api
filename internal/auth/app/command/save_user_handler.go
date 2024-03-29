package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/hasher"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// SaveUserHandler represents the SaveUser command.
type SaveUserHandler = command.Handler[*SaveUserInput]

// saveUserHandlerImpl represents the implementation of the SaveUser command.
type saveUserHandlerImpl struct {
	repo.User
	hasher.Hasher
	publisher message.Publisher[user.User]
}

// NewSaveUserHandler returns a new Save instance.
func NewSaveUserHandler(
	userRepo repo.User,
	hasher hasher.Hasher,
	publisher message.Publisher[user.User],
) SaveUserHandler {
	errutil.MustBeNotEmpty("repo.User", userRepo)
	errutil.MustBeNotEmpty("hasher.Hasher", hasher)
	errutil.MustBeNotEmpty("message.Publisher", publisher)

	return &saveUserHandlerImpl{userRepo, hasher, publisher}
}

// Handle implements command.Handler.
func (c *saveUserHandlerImpl) Handle(ctx context.Context, input *SaveUserInput) error {
	hashedPassword, err := c.Hash(input.Password)
	if err != nil {
		return err
	}

	u, err := user.NewUser().
		SetEmail(input.Email).
		SetName(input.Name).
		SetPassword(hashedPassword).
		Validate()
	if err != nil {
		return err
	}

	u, err = c.Save(ctx, repo.SaveUserInput{User: u})
	if err != nil {
		return err
	}

	c.publisher.Handle(*u)

	return nil
}
