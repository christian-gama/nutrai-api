package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/internal/exception/app/event"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		rabbitmq.MakeConsumer(
			"exceptions",
			event.SaveException.String(),
			event.SaveException.String(),
		),
		persistence.MakeSQLException(),
	)
}
