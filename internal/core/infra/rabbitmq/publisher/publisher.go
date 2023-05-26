package publisher

import (
	"context"
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	amqp "github.com/rabbitmq/amqp091-go"
)

// publisherImpl is the implementation of a message publisher.
type publisherImpl struct {
	*rabbitmq.RabbitMQ
	options *options
}

// NewPublisher creates a new message publisher.
func NewPublisher(rmq *rabbitmq.RabbitMQ, opts ...func(*options)) message.Publisher {
	options := &options{
		ExchangeName: "",
		RoutingKey:   "",
		ContentType:  "text/plain",
		Kind:         amqp.ExchangeDirect,
		Durable:      false,
		AutoDelete:   false,
		Internal:     false,
		NoWait:       false,
		Args:         nil,
		Mandatory:    false,
		Immediate:    false,
	}
	for _, opt := range opts {
		opt(options)
	}

	errutil.MustBeNotEmpty("exchange name", options.ExchangeName)
	errutil.MustBeNotEmpty("routing key", options.RoutingKey)

	return &publisherImpl{rmq, options}
}

// Handle publishes a message to the broker.
func (p *publisherImpl) Handle(msg []byte) {
	ch := p.ChannelPool()
	defer p.ReleaseChannelPool(ch)

	err := ch.ExchangeDeclare(
		p.options.ExchangeName,
		p.options.Kind,
		p.options.Durable,
		p.options.AutoDelete,
		p.options.Internal,
		p.options.NoWait,
		p.options.Args,
	)
	if err != nil {
		panic(errors.InternalServerError("could not declare an exchange"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		p.options.ExchangeName,
		p.options.RoutingKey,
		p.options.Mandatory,
		p.options.Immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		panic(errors.InternalServerError(fmt.Sprintf("could not publish message: %s", err)))
	}
}
