package service

import value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"

// RefreshTokenInput is the input of the RefreshToken service.
type RefreshTokenInput struct {
	Refresh value.Token `json:"refresh" validate:"required,jwt" faker:"jwt"`
}

// RefreshTokenOutput is the output of the RefreshToken service.
type RefreshTokenOutput struct {
	Access value.Token `json:"access" faker:"jwt"`
}
