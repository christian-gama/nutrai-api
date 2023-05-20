package command

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/hasher"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// CheckCredentialsHandler is the handler for CheckCredentials.
type CheckCredentialsHandler = command.Handler[*CheckCredentialsInput]

// checkCredentialsHandlerImpl is the implementation of CheckCredentialsHandler.
type checkCredentialsHandlerImpl struct {
	repo.User
	hasher.Hasher
}

// NewCheckCredentialsHandler returns a new CheckCredentialsHandler.
func NewCheckCredentialsHandler(userRepo repo.User, hasher hasher.Hasher) CheckCredentialsHandler {
	if userRepo == nil {
		panic(errors.New("repo.User cannot be nil"))
	}

	if hasher == nil {
		panic(errors.New("hasher.Hasher cannot be nil"))
	}

	return &checkCredentialsHandlerImpl{userRepo, hasher}
}

// Handle implements command.Handler.
func (c *checkCredentialsHandlerImpl) Handle(
	ctx context.Context,
	input *CheckCredentialsInput,
) error {
	user, err := c.FindByEmail(ctx, repo.FindByEmailUserInput{Email: input.Email})
	if err != nil {
		return err
	}

	password := input.Password
	if err := password.Validate(); err != nil {
		return err
	}

	if err := c.checkPassword(user, password); err != nil {
		return err
	}

	return nil
}

// checkPassword checks if the given password matches the user's hashed password.
func (c *checkCredentialsHandlerImpl) checkPassword(
	user *user.User,
	password value.Password,
) error {
	if err := c.Hasher.Compare(password, user.Password); err != nil {
		return errutil.NewErrInvalid("password", "does not match")
	}
	return nil
}
