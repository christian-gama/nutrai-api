package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		publisher.MakePublisher(
			publisher.WithExchangeName(event.SaveException),
			publisher.WithRoutingKey(event.SaveException),
		),
	)
}
