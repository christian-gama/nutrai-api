package service

import "github.com/christian-gama/nutrai-api/internal/auth/domain/value"

type LoginInput struct {
	Email    value.Email    `json:"email" validate:"required,email" faker:"email"`
	Password value.Password `json:"password" validate:"required,min=8,max=32" faker:"len=8"`
}

type LoginOutput struct {
	Access  value.Token `json:"access"`
	Refresh value.Token `json:"refresh"`
}
