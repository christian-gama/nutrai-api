package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
)

func MakeSaveException() SaveException {
	return NewSaveException(command.MakeSaveExceptionHandler())
}
