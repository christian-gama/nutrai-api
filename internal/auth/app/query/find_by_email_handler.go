package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/repo"
	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// FindByEmailInput is the query to find a user by email.
type FindByEmailHandler = query.Handler[*FindByEmailInput, *FindByEmailOutput]

// findByEmailHandlerImpl is the implementation of the FindByEmailHandler interface.
type findByEmailHandlerImpl struct {
	repo.User
}

// NewFindByEmailHandler creates a new instance of the FindByEmailHandler interface.
func NewFindByEmailHandler(userRepo repo.User) FindByEmailHandler {
	errutil.MustBeNotEmpty("repo.User", userRepo)

	return &findByEmailHandlerImpl{
		User: userRepo,
	}
}

// Handle implements the FindByEmailHandler interface.
func (q *findByEmailHandlerImpl) Handle(
	ctx context.Context,
	input *FindByEmailInput,
) (*FindByEmailOutput, error) {
	user, err := q.FindByEmail(ctx, repo.FindByEmailUserInput{Email: input.Email})
	if err != nil {
		return nil, err
	}

	return &FindByEmailOutput{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
