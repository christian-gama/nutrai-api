package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/internal/exception/app/event"
)

func MakeSaveExceptionHandler() SaveExceptionHandler {
	return NewSaveExceptionHandler(
		publisher.MakePublisher(
			publisher.WithExchange("exceptions"),
			publisher.WithRoutingKey(event.SaveException),
		),
	)
}
