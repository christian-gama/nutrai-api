package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
	"github.com/christian-gama/nutrai-api/internal/auth/event"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
)

func MakeSendWelcomeHandler() SendWelcomeHandler {
	return NewSendWelcomeHandler(
		consumer.MakeConsumer[command.SaveUserInput](
			consumer.WithExchangeName(event.User),
			consumer.WithRoutingKey(event.SaveUser),
		),
		mailer.MakeMailer(),
	)
}
