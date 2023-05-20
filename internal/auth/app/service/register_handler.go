package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/jwt"
	"github.com/christian-gama/nutrai-api/internal/core/app/service"
	userCmd "github.com/christian-gama/nutrai-api/internal/user/app/command"
)

// RegisterHandler handles the register request and returns the access and refresh tokens.
type RegisterHandler = service.Handler[*RegisterInput, *RegisterOutput]

type registerHandlerImpl struct {
	accessToken        jwt.Generator
	refreshToken       jwt.Generator
	savePatientHandler userCmd.SavePatientHandler
}

// NewRegisterHandler creates a new RegisterHandler instance.
func NewRegisterHandler(
	accessToken jwt.Generator,
	refreshToken jwt.Generator,
	savePatientHandler userCmd.SavePatientHandler,
) RegisterHandler {
	return &registerHandlerImpl{
		accessToken:        accessToken,
		refreshToken:       refreshToken,
		savePatientHandler: savePatientHandler,
	}
}

// Handle implements the RegisterHandler interface.
func (h *registerHandlerImpl) Handle(
	ctx context.Context,
	input *RegisterInput,
) (*RegisterOutput, error) {
	if err := h.savePatientHandler.Handle(ctx, &userCmd.SavePatientInput{
		Age:      input.Age,
		HeightM:  input.HeightM,
		WeightKG: input.WeightKG,
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
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

	return &RegisterOutput{
		Access:  access,
		Refresh: refresh,
	}, nil
}
