package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// LogoutHandler represents the Logout command.
type LogoutHandler = command.Handler[*LogoutInput]

// logoutHandlerImpl represents the implementation of the Logout command.
type logoutHandlerImpl struct {
	repo.Token
	jwt.Verifier
}

// NewLogoutHandler returns a new Delete instance.
func NewLogoutHandler(tokenRepo repo.Token, verifier jwt.Verifier) LogoutHandler {
	errutil.MustBeNotEmpty("repo.Token", tokenRepo)
	errutil.MustBeNotEmpty("jwt.Verifier (Refresh)", verifier)

	return &logoutHandlerImpl{tokenRepo, verifier}
}

// Handle implements command.Handler.
func (c *logoutHandlerImpl) Handle(ctx context.Context, input *LogoutInput) error {
	claims, err := c.Verify(input.Refresh, true)
	if err != nil {
		return errors.Unauthorized(err.Error())
	}

	if claims.Sub.Email != input.User.Email {
		return errors.Unauthorized("invalid token")
	}

	t, err := c.Find(ctx, repo.FindTokenInput{Email: input.User.Email, Jti: claims.Jti})
	if err != nil {
		return errors.Unauthorized(err.Error())
	}

	err = c.Delete(ctx, repo.DeleteTokenInput{Email: input.User.Email, Jti: t.Jti})
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
