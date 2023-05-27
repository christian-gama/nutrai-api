package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// LogoutHandler represents the Logout command.
type LogoutHandler = command.Handler[*LogoutInput]

// logoutHandlerImpl represents the implementation of the Logout command.
type logoutHandlerImpl struct {
	repo.Token
}

// NewLogoutHandler returns a new Delete instance.
func NewLogoutHandler(tokenRepo repo.Token) LogoutHandler {
	errutil.MustBeNotEmpty("repo.Token", tokenRepo)

	return &logoutHandlerImpl{tokenRepo}
}

// Handle implements command.Handler.
func (c *logoutHandlerImpl) Handle(ctx context.Context, input *LogoutInput) error {
	_, err := c.Find(ctx, repo.FindTokenInput{Email: input.User.Email})
	if err != nil {
		return err
	}

	return c.Delete(ctx, repo.DeleteTokenInput{Email: input.User.Email})
}
