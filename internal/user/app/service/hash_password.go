package service

import value "github.com/christian-gama/nutrai-api/internal/user/domain/value/user"

// HashPasswordInput represents the input data for the HashPassword service.
type HashPasswordInput struct {
	Password value.Password `json:"password" faker:"len=8"`
}

// HashPasswordOutput represents the output data for the HashPassword service.
type HashPasswordOutput struct {
	Password value.Password `json:"hashedPassword"`
}
