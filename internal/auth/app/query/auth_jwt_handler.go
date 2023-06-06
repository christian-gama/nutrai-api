package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// AuthInput is the query to find a user by email.
type AuthJwtHandler = query.Handler[*AuthJwtInput, *AuthJwtOutput]

// AuthJwtHandlerImpl is the implementation of the AuthHandler interface.
type AuthJwtHandlerImpl struct {
	repo.User
	jwt.Verifier
}

// NewAuthJwtHandler creates a new instance of the AuthHandler interface.
func NewAuthJwtHandler(userRepo repo.User, verifier jwt.Verifier) AuthJwtHandler {
	errutil.MustBeNotEmpty("repo.User", userRepo)
	errutil.MustBeNotEmpty("jwt.Verifier (Access)", verifier)

	return &AuthJwtHandlerImpl{
		User:     userRepo,
		Verifier: verifier,
	}
}

// Handle implements the AuthHandler interface.
func (q *AuthJwtHandlerImpl) Handle(
	ctx context.Context,
	input *AuthJwtInput,
) (*AuthJwtOutput, error) {
	claims, err := q.Verify(input.Access, false)
	if err != nil {
		return nil, errors.Unauthorized(err.Error())
	}

	user, err := q.FindByEmail(ctx, repo.FindByEmailUserInput{Email: claims.Sub.Email})
	if err != nil {
		return nil, errors.Unauthorized("you are not authorized to access this resource")
	}

	return &AuthJwtOutput{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}, nil
}
