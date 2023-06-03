package command

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
)

func MakeRecoveryHandler() RecoveryHandler {
	return NewRecoveryHandler(
		publisher.MakePublisher[RecoveryInput](
			publisher.WithExchangeName(event.Exception),
			publisher.WithRoutingKey(event.Recovery),
		),
	)
}
