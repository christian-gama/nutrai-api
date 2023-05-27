package service

import (
	"context"

	userCmd "github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/app/service"
)

// RegisterHandler handles the register request and returns the access and refresh tokens.
type RegisterHandler = service.Handler[*RegisterInput, *RegisterOutput]

// registerHandlerImpl is the implementation of the RegisterHandler interface.
type registerHandlerImpl struct {
	accessToken     jwt.Generator
	refreshToken    jwt.Generator
	saveUserHandler userCmd.SaveUserHandler
}

// NewRegisterHandler creates a new RegisterHandler instance.
func NewRegisterHandler(
	accessToken jwt.Generator,
	refreshToken jwt.Generator,
	saveUserHandler userCmd.SaveUserHandler,
) RegisterHandler {
	return &registerHandlerImpl{
		accessToken:     accessToken,
		refreshToken:    refreshToken,
		saveUserHandler: saveUserHandler,
	}
}

// Handle implements the RegisterHandler interface.
func (h *registerHandlerImpl) Handle(
	ctx context.Context,
	input *RegisterInput,
) (*RegisterOutput, error) {
	if err := h.saveUserHandler.Handle(ctx, &userCmd.SaveUserInput{
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
	}); err != nil {
		return nil, err
	}

	subject := &jwt.Subject{Email: input.Email}
	accessToken, err := h.accessToken.Generate(subject, false)
	if err != nil {
		return nil, err
	}

	refreshToken, err := h.refreshToken.Generate(subject, true)
	if err != nil {
		return nil, err
	}

	return &RegisterOutput{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}
