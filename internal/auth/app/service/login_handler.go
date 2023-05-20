package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/app/service"
	userCmd "github.com/christian-gama/nutrai-api/internal/user/app/command"
)

type LoginHandler = service.Handler[*LoginInput, *LoginOutput]

type loginHandlerImpl struct {
	accessToken             jwt.Generator
	refreshToken            jwt.Generator
	checkCredentialsHandler userCmd.CheckCredentialsHandler
}

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

func (h *loginHandlerImpl) Handle(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	if err := h.checkCredentialsHandler.Handle(ctx, &userCmd.CheckCredentialsInput{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	access, err := h.accessToken.Generate(&jwt.Subject{Email: input.Email})
	if err != nil {
		return nil, err
	}

	refresh, err := h.refreshToken.Generate(&jwt.Subject{Email: input.Email})
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Access:  access,
		Refresh: refresh,
	}, nil
}
