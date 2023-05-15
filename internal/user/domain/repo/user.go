package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/queryer"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// SaveUserInput is the input for the Save method.
type SaveUserInput struct {
	User *user.User
}

// AllUsersInput is the input for the All method.
type AllUsersInput struct {
	queryer.Filterer
	queryer.Sorter
	queryer.Paginator
	queryer.Preloader
}

// FindUserInput is the input for the Find method.
type FindUserInput struct {
	ID coreValue.ID
	queryer.Filterer
	queryer.Preloader
}

// FindByEmailUserInput is the input for the FindByEmail method.
type FindByEmailUserInput struct {
	Email value.Email
}

// DeleteUserInput is the input for the Delete method.
type DeleteUserInput struct {
	IDs []coreValue.ID
}

// UpdateUserInput is the input for the Update method.
type UpdateUserInput struct {
	User *user.User
	ID   coreValue.ID
}

// User is the interface that wraps the basic User methods.
type User interface {
	// All returns all users.
	All(ctx context.Context, input AllUsersInput) (*queryer.PaginationOutput[*user.User], error)

	// Delete deletes the user with the given id.
	Delete(ctx context.Context, input DeleteUserInput) error

	// Find returns the user with the given id.
	Find(ctx context.Context, input FindUserInput) (*user.User, error)

	// FindByEmail returns the user with the given email.
	FindByEmail(ctx context.Context, input FindByEmailUserInput) (*user.User, error)

	// Save saves the given user.
	Save(ctx context.Context, input SaveUserInput) (*user.User, error)

	// Update updates the given user.
	Update(ctx context.Context, input UpdateUserInput) error
}
