package rabbitmq

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
)

var (
	consumerConnection  *RabbitMQ
	publisherConnection *RabbitMQ
)

func MakePublisher(exchange string, routingKey string) message.Publisher {
	if publisherConnection == nil {
		publisherConnection = NewConnection(log.WithCaller, "publisher")
	}

	return NewPublisher(publisherConnection, exchange, routingKey)
}

func MakeConsumer(exchange string, queue string, routingKey string) message.Consumer {
	if consumerConnection == nil {
		consumerConnection = NewConnection(log.WithCaller, "consumer")
	}

	return NewConsumer(consumerConnection, exchange, routingKey, queue, log.WithCaller)
}
