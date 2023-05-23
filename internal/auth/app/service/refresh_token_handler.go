package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/app/service"
)

// RefreshTokenHandler is the interface that wraps the Handle method.
type RefreshTokenHandler = service.Handler[*RefreshTokenInput, *RefreshTokenOutput]

// refreshTokenHandlerImpl is the implementation of the RefreshTokenHandler interface.
type refreshTokenHandlerImpl struct {
	accessTokenGenerator jwt.Generator
	verifier             jwt.Verifier
}

// NewRefreshTokenHandler returns a new RefreshTokenHandler.
func NewRefreshTokenHandler(
	accessTokenGenerator jwt.Generator,
	verifier jwt.Verifier,
) RefreshTokenHandler {
	return &refreshTokenHandlerImpl{
		accessTokenGenerator: accessTokenGenerator,
		verifier:             verifier,
	}
}

// Handle implements the RefreshTokenHandler interface.
func (h *refreshTokenHandlerImpl) Handle(
	ctx context.Context,
	input *RefreshTokenInput,
) (*RefreshTokenOutput, error) {
	claims, err := h.verifier.Verify(input.Refresh)
	if err != nil {
		return nil, err
	}

	accessToken, err := h.accessTokenGenerator.Generate(&claims.Sub)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenOutput{
		Access: accessToken,
	}, nil
}
