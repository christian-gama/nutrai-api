package repo

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/token"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

// SaveTokenInput is the input for the Save method.
type SaveTokenInput struct {
	Token *token.Token
}

// FindTokenInput is the input for the Find method.
type FindTokenInput struct {
	Email value.Email
}

// DeleteTokenInput is the input for the Delete method.
type DeleteTokenInput struct {
	Email value.Email
}

// Token is the interface that wraps the basic Token methods.
type Token interface {
	// Delete deletes the token with the given id.
	Delete(ctx context.Context, input DeleteTokenInput) error

	// Find returns the token with the given id.
	Find(ctx context.Context, input FindTokenInput) (*token.Token, error)

	// Save saves the given token.
	Save(ctx context.Context, input SaveTokenInput) (*token.Token, error)
}
