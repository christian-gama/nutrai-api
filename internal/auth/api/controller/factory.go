package controller

import "github.com/christian-gama/nutrai-api/internal/auth/app/service"

func MakeLogin() Login {
	return NewLogin(service.MakeLoginHandler())
}

func MakeRegister() Register {
	return NewRegister(service.MakeRegisterHandler())
}
