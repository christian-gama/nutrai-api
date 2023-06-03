package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// LogoutAllHandler represents the LogoutAll command.
type LogoutAllHandler = command.Handler[*LogoutAllInput]

// logoutAllHandlerImpl represents the implementation of the LogoutAll command.
type logoutAllHandlerImpl struct {
	repo.Token
}

// NewLogoutAllHandler returns a new Delete instance.
func NewLogoutAllHandler(tokenRepo repo.Token) LogoutAllHandler {
	errutil.MustBeNotEmpty("repo.Token", tokenRepo)

	return &logoutAllHandlerImpl{tokenRepo}
}

// Handle implements command.Handler.
func (c *logoutAllHandlerImpl) Handle(ctx context.Context, input *LogoutAllInput) error {
	err := c.DeleteAll(ctx, repo.DeleteAllTokenInput{Email: input.User.Email})
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
