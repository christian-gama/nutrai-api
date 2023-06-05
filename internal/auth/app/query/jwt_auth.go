package query

import (
	jwtValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	userValue "github.com/christian-gama/nutrai-api/internal/auth/domain/value/user"
	coreValue "github.com/christian-gama/nutrai-api/internal/core/domain/value"
)

// JwtAuthInput is the query to check if the JWT token is valid and find the user associated with
// it.
type JwtAuthInput struct {
	Access jwtValue.Token `json:"token" validate:"required,jwt" faker:"jwt"`
}

// JwtAuthOutput is the output of the AuthInput query.
type JwtAuthOutput struct {
	ID       coreValue.ID       `json:"id" faker:"uint"`
	Email    userValue.Email    `json:"email" faker:"email"`
	Name     userValue.Name     `json:"name" faker:"name"`
	Password userValue.Password `json:"-" faker:"len=8"`
}
