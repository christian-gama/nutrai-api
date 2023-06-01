package consumer

import (
	"github.com/christian-gama/nutrai-api/config/env"
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
			consumer.WithDelayBetweenMessages(env.Mailer.DelayBetweenEmails),
		),
		mailer.MakeMailer(),
	)
}
