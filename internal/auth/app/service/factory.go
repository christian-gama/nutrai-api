package service

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
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
		command.MakeSaveUserHandler(),
	)
}

func MakeRefreshTokenHandler() RefreshTokenHandler {
	return NewRefreshTokenHandler(
		jwt.MakeAccessTokenGenerator(),
		jwt.MakeVerifier(),
	)
}
