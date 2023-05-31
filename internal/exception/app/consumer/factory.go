package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		consumer.MakeConsumer[exception.Exception](
			consumer.WithExchangeName(event.Exception),
			consumer.WithRoutingKey(event.CatchException),
		),
		persistence.MakeSQLException(),
	)
}
