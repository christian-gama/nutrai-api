package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/command"
	"github.com/christian-gama/nutrai-api/internal/user/domain/hasher"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/user/domain/repo"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
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
func NewCheckCredentialsHandler(r repo.User, h hasher.Hasher) CheckCredentialsHandler {
	return &checkCredentialsHandlerImpl{r, h}
}

// Handle implements command.Handler.
func (c *checkCredentialsHandlerImpl) Handle(ctx context.Context, input *CheckCredentialsInput) error {
	user, err := c.FindByEmail(ctx, repo.FindByEmailUserInput{Email: input.Email})
	if err != nil {
		return err
	}

	if err := c.checkPassword(user, value.Password(input.Password)); err != nil {
		return err
	}

	return nil
}

// checkPassword checks if the given password matches the user's hashed password.
func (c *checkCredentialsHandlerImpl) checkPassword(user *user.User, password value.Password) error {
	if err := c.Hasher.Compare(password, user.Password); err != nil {
		return errutil.NewErrInvalid("password", "does not match")
	}
	return nil
}
