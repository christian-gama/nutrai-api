package controller

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
)

func MakeLogin() Login {
	return NewLogin(service.MakeLoginHandler())
}

func MakeRegister() Register {
	return NewRegister(service.MakeRegisterHandler())
}

func MakeDeleteUser() DeleteMe {
	return NewDeleteMe(command.MakeDeleteMeHandler())
}

func MakeRefreshToken() RefreshToken {
	return NewRefreshToken(service.MakeRefreshTokenHandler())
}

func MakeChangePassword() ChangePassword {
	return NewChangePassword(command.MakeChangePasswordHandler())
}

func MakeLogout() Logout {
	return NewLogout(command.MakeLogoutHandler())
}

func MakeLogoutAll() LogoutAll {
	return NewLogoutAll(command.MakeLogoutAllHandler())
}
