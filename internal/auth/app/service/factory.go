package service

import (
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
	"github.com/christian-gama/nutrai-api/internal/user/app/command"
)

func MakeLoginHandler() LoginHandler {
	return NewLoginHandler(
		jwt.MakeAccessTokenGenerator(),
		jwt.MakeRefreshTokenGenerator(),
		command.MakeCheckCredentialsHandler(),
	)
}

func MakeRegisterHandler() RegisterHandler {
	return NewRegisterHandler(
		jwt.MakeAccessTokenGenerator(),
		jwt.MakeRefreshTokenGenerator(),
		command.MakeSavePatientHandler(),
	)
}
