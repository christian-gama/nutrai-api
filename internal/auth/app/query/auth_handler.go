package query

import (
	"context"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/app/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// AuthInput is the query to find a user by email.
type AuthHandler = query.Handler[*AuthInput, *AuthOutput]

// checkTokenHandlerImpl is the implementation of the AuthHandler interface.
type checkTokenHandlerImpl struct {
	repo.User
	jwt.Verifier
}

// NewAuthHandler creates a new instance of the AuthHandler interface.
func NewAuthHandler(userRepo repo.User, verifier jwt.Verifier) AuthHandler {
	if userRepo == nil {
		panic(errors.New("repo.User is required"))
	}

	if verifier == nil {
		panic(errors.New("jwt.Verifier is required"))
	}

	return &checkTokenHandlerImpl{
		User:     userRepo,
		Verifier: verifier,
	}
}

// Handle implements the AuthHandler interface.
func (q *checkTokenHandlerImpl) Handle(
	ctx context.Context,
	input *AuthInput,
) (*AuthOutput, error) {
	claims, err := q.Verify(input.Access)
	if err != nil {
		return nil, errutil.Unauthorized(err.Error())
	}

	user, err := q.FindByEmail(ctx, repo.FindByEmailUserInput{Email: claims.Sub.Email})
	if err != nil {
		return nil, errutil.Unauthorized("you are not authorized to access this resource")
	}

	return &AuthOutput{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}, nil
}
