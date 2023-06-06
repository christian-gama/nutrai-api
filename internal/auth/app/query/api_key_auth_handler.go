package query

import (
	"context"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/domain/query"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// AuthInput is the query to find a user by email.
type ApiKeyAuthHandler = query.Handler[*ApiKeyAuthInput, *ApiKeyAuthOutput]

// apiKeyAuthHandlerImpl is the implementation of the AuthHandler interface.
type apiKeyAuthHandlerImpl struct{}

// NewApiKeyAuthHandler creates a new instance of the AuthHandler interface.
func NewApiKeyAuthHandler() ApiKeyAuthHandler {
	return &apiKeyAuthHandlerImpl{}
}

// Handle implements the AuthHandler interface.
func (q *apiKeyAuthHandlerImpl) Handle(
	ctx context.Context,
	input *ApiKeyAuthInput,
) (*ApiKeyAuthOutput, error) {
	if input.Key != env.App.ApiKey {
		return nil, errors.Unauthorized("invalid api key")
	}

	return &ApiKeyAuthOutput{
		Key: input.Key,
	}, nil
}
