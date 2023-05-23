package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/app/event"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		consumer.MakeConsumer(
			consumer.WithExchange("exceptions"),
			consumer.WithRoutingKey(event.SaveException),
		),
		persistence.MakeSQLException(),
	)
}
