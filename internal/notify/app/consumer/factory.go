package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/event"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
)

func MakeSendWelcomeHandler() SendWelcomeHandler {
	return NewSendWelcomeHandler(
		consumer.MakeConsumer[user.User](
			consumer.WithExchangeName(event.User),
			consumer.WithRoutingKey(event.SaveUser),
		),
		mailer.MakeMailer(),
	)
}
