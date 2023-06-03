package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/internal/exception/event"
	persistence "github.com/christian-gama/nutrai-api/internal/exception/infra/persistence/sql"
)

func MakeRecoveryHandler() RecoveryHandler {
	return NewRecoveryHandler(
		consumer.MakeConsumer[command.RecoveryInput](
			consumer.WithExchangeName(event.Exception),
			consumer.WithRoutingKey(event.Recovery),
		),
		persistence.MakeSQLException(),
	)
}
