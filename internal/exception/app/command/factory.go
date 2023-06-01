package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
)

func MakeCatchExceptionHandler() CatchExceptionHandler {
	return NewCatchExceptionHandler(
		publisher.MakePublisher[CatchExceptionInput](
			publisher.WithExchangeName(event.Exception),
			publisher.WithRoutingKey(event.CatchException),
		),
	)
}
