package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/internal/exception/app/event"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		rabbitmq.MakePublisher("exceptions", event.SaveException.String()),
	)
}
