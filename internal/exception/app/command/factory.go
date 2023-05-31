package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
)

func MakeCatchExceptionHandler() CatchExceptionHandler {
	return NewCatchExceptionHandler(
		publisher.MakePublisher[exception.Exception](
			publisher.WithExchangeName(event.Exception),
			publisher.WithRoutingKey(event.CatchException),
		),
	)
}
