package service

import (
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

type LoginInput struct {
	Email    userValue.Email    `json:"email" validate:"required,email" faker:"email"`
	Password userValue.Password `json:"password" validate:"required,user_password" faker:"len=8"`
}

type LoginOutput struct {
	Access  jwtValue.Token `json:"access"`
	Refresh jwtValue.Token `json:"refresh"`
}
