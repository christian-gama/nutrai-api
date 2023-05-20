package service

import (
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

type RegisterInput struct {
	Email    userValue.Email    `json:"email" validate:"required,email" faker:"email"`
	Name     userValue.Name     `json:"name" validate:"required,min=2,max=100" faker:"name"`
	Password userValue.Password `json:"password" validate:"required,min=8,max=32" faker:"len=8"`
}

type RegisterOutput struct {
	Access  jwtValue.Token `json:"access"`
	Refresh jwtValue.Token `json:"refresh"`
}
