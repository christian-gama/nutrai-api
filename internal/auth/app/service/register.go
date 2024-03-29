package service

import (
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
)

// RegisterInput is the input of the Register service.
type RegisterInput struct {
	Email    userValue.Email    `json:"email" validate:"required,email" faker:"email"`
	Name     userValue.Name     `json:"name" validate:"required,user_name" faker:"name"`
	Password userValue.Password `json:"password" validate:"required,user_password" faker:"len=8"`
}

// RegisterOutput is the output of the Register service.
type RegisterOutput struct {
	Access  jwtValue.Token `json:"access" faker:"jwt"`
	Refresh jwtValue.Token `json:"refresh" faker:"jwt"`
}
