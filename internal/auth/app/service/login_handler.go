package service

import (
	"context"

	userCmd "github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/app/service"
)

// LoginHandler handles the login request and returns the access and refresh tokens.
type LoginHandler = service.Handler[*LoginInput, *LoginOutput]

type loginHandlerImpl struct {
	accessToken             jwt.Generator
	refreshToken            jwt.Generator
	checkCredentialsHandler userCmd.CheckCredentialsHandler
}

// NewLoginHandler creates a new LoginHandler instance.
func NewLoginHandler(
	accessToken jwt.Generator,
	refreshToken jwt.Generator,
	checkCredentialsHandler userCmd.CheckCredentialsHandler,
) LoginHandler {
	return &loginHandlerImpl{
		accessToken:             accessToken,
		refreshToken:            refreshToken,
		checkCredentialsHandler: checkCredentialsHandler,
	}
}

// Handle implements the LoginHandler interface.
func (h *loginHandlerImpl) Handle(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	if err := h.checkCredentialsHandler.Handle(ctx, &userCmd.CheckCredentialsInput{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	subject := &jwt.Subject{Email: input.Email}
	access, err := h.accessToken.Generate(subject)
	if err != nil {
		return nil, err
	}

	refresh, err := h.refreshToken.Generate(subject)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Access:  access,
		Refresh: refresh,
	}, nil
}
