package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	amqp "github.com/rabbitmq/amqp091-go"
)

// publisherImpl is the implementation of a message publisher.
type publisherImpl struct {
	*RabbitMQ
	exchange   string
	routingKey string
}

// NewPublisher creates a new message publisher.
func NewPublisher(conn *RabbitMQ, exchange string, routingKey string) message.Publisher {
	return &publisherImpl{conn, exchange, routingKey}
}

// Handle publishes a message to the broker.
func (p *publisherImpl) Handle(msg []byte) {
	ch := p.ChannelPool()
	defer p.ReleaseChannelPool(ch)

	err := ch.ExchangeDeclare(
		p.exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(errutil.NewErrInternal("could not declare an exchange"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		p.exchange,
		p.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		panic(errutil.NewErrInternal(fmt.Sprintf("could not publish message: %s", err)))
	}
}
