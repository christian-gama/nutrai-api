package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/christian-gama/nutrai-api/internal/shared/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
)

// SaveUserInput is the input for the Save method.
type SaveUserInput struct {
	User *user.User
}

// AllUsersInput is the input for the All method.
type AllUsersInput struct {
	querying.Filterer
	querying.Sorter
	querying.Paginator
}

// FindUserInput is the input for the Find method.
type FindUserInput struct {
	ID value.ID
	querying.Filterer
}

// FindByEmailUserInput is the input for the FindByEmail method.
type FindByEmailUserInput struct {
	Email string
}

// DeleteUserInput is the input for the Delete method.
type DeleteUserInput struct {
	IDs []value.ID
}

// UpdateUserInput is the input for the Update method.
type UpdateUserInput struct {
	User *user.User
	ID   value.ID
}

// User is the interface that wraps the basic User methods.
type User interface {
	// All returns all users.
	All(ctx context.Context, input AllUsersInput, preload ...string) (*querying.PaginationOutput[*user.User], error)

	// Delete deletes the user with the given id.
	Delete(ctx context.Context, input DeleteUserInput) error

	// Find returns the user with the given id.
	Find(ctx context.Context, input FindUserInput, preload ...string) (*user.User, error)

	// FindByEmail returns the user with the given email.
	FindByEmail(ctx context.Context, input FindByEmailUserInput) (*user.User, error)

	// Save saves the given user.
	Save(ctx context.Context, input SaveUserInput) (*user.User, error)

	// Update updates the given user.
	Update(ctx context.Context, input UpdateUserInput) error
}
