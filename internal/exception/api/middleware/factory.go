package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(command.MakeSaveExceptionHandler())
}
