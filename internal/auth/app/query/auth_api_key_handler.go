package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// AuthInput is the query to find a user by email.
type AuthApiKeyHandler = query.Handler[*AuthApiKeyInput, *AuthApiKeyOutput]

// AuthApiKeyHandlerImpl is the implementation of the AuthHandler interface.
type AuthApiKeyHandlerImpl struct{}

// NewAuthApiKeyHandler creates a new instance of the AuthHandler interface.
func NewAuthApiKeyHandler() AuthApiKeyHandler {
	return &AuthApiKeyHandlerImpl{}
}

// Handle implements the AuthHandler interface.
func (q *AuthApiKeyHandlerImpl) Handle(
	ctx context.Context,
	input *AuthApiKeyInput,
) (*AuthApiKeyOutput, error) {
	if input.Key != env.App.ApiKey {
		return nil, errors.Unauthorized("invalid api key")
	}

	return &AuthApiKeyOutput{
		Key: input.Key,
	}, nil
}
