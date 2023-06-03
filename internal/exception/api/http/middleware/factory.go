package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
)

func MakeRecovery() Recovery {
	return NewRecovery(command.MakeRecoveryHandler())
}
