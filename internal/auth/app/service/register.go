package service

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/value"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
)

type RegisterInput = command.SavePatientInput

type RegisterOutput struct {
	Access  value.Token `json:"access"`
	Refresh value.Token `json:"refresh"`
}
