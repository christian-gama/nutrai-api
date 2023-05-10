package service

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/shared/app/service"
	"github.com/christian-gama/nutrai-api/internal/user/domain/hasher"
	value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"
)

// HashPasswordInput represents the input data for the HashPassword service.
type HashPasswordHandler = service.Handler[*HashPasswordInput, *HashPasswordOutput]

// hashPasswordHandlerImpl represents the implementation of the HashPassword service.
type hashPasswordHandlerImpl struct {
	hasher.Hasher
}

// NewHashPasswordHandler creates a new instance of a HashPassword service.
func NewHashPasswordHandler(h hasher.Hasher) HashPasswordHandler {
	return &hashPasswordHandlerImpl{h}
}

// Handle generates a hash for the given password.
func (s *hashPasswordHandlerImpl) Handle(ctx context.Context, input *HashPasswordInput) (*HashPasswordOutput, error) {
	hashedPassword, err := s.Hash(value.Password(input.Password))
	if err != nil {
		return nil, err
	}

	return &HashPasswordOutput{hashedPassword}, nil
}
