package consumer

import (
	"encoding/json"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	amqp "github.com/rabbitmq/amqp091-go"
)

// consumerImpl is a RabbitMQ consumer implementation.
type consumerImpl[Data any] struct {
	rmq     *rabbitmq.RabbitMQ
	log     logger.Logger
	options *options
}

// NewConsumer creates a new RabbitMQ consumer.
func NewConsumer[Data any](
	rmq *rabbitmq.RabbitMQ,
	log logger.Logger,
	opts ...func(*options),
) message.Consumer[Data] {
	options := &options{
		ExchangeName:    "",
		RoutingKey:      "",
		Kind:            rabbitmq.ExchangeDirect,
		Durable:         false,
		AutoDelete:      false,
		Internal:        false,
		NoWait:          false,
		Args:            nil,
		QueueName:       "",
		QueueDurable:    false,
		QueueAutoDelete: false,
		QueueExclusive:  false,
		QueueNoWait:     false,
		QueueArgs:       nil,
	}
	for _, opt := range opts {
		opt(options)
	}

	errutil.MustBeNotEmpty("options.ExchangeName", options.ExchangeName)
	errutil.MustBeNotEmpty("options.RoutingKey", options.RoutingKey)

	if options.QueueName == "" {
		options.QueueName = options.RoutingKey
	}

	return &consumerImpl[Data]{rmq, log, options}
}

// Handle handles a message.
func (c *consumerImpl[Data]) Handle(handler func(data Data) error) {
	ch, err := c.rmq.Connection().Channel()
	if err != nil {
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		c.options.ExchangeName,
		c.options.Kind,
		c.options.Durable,
		c.options.AutoDelete,
		c.options.Internal,
		c.options.NoWait,
		c.options.Args,
	)
	if err != nil {
		return
	}

	q, err := ch.QueueDeclare(
		c.options.QueueName,
		c.options.QueueDurable,
		c.options.QueueAutoDelete,
		c.options.QueueExclusive,
		c.options.QueueNoWait,
		c.options.QueueArgs,
	)
	if err != nil {
		return
	}

	if err := ch.Qos(1, 0, false); err != nil {
		return
	}

	err = ch.QueueBind(
		q.Name,
		c.options.RoutingKey,
		c.options.ExchangeName,
		c.options.NoWait,
		c.options.Args,
	)
	if err != nil {
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			c.handle(msg, handler)
		}
	}()

	<-forever
}

func (c *consumerImpl[Data]) handle(msg amqp.Delivery, handler func(data Data) error) {
	defer func() {
		if r := recover(); r != nil {
			c.log.Warnf("Consumer | %12s | Recovered from panic: %s", c.options.QueueName, r)
			msg.Nack(false, false)
		}
	}()

	var err error
	duration := bench.Duration(func() {
		var data Data
		err = json.Unmarshal(msg.Body, &data)
		if err != nil {
			c.log.Warnf(
				"Consumer | %s | Could not unmarshal message: %s",
				c.options.QueueName,
				err.Error(),
			)
			msg.Nack(false, false)
			return
		}

		err = handler(data)
	})

	if err != nil {
		c.log.Warnf(
			"Consumer | %s | Could not process message: %s",
			c.options.QueueName,
			err.Error(),
		)
	} else {
		c.log.Infof("Consumer | %12s | Message processed successfully in %s", c.options.QueueName, duration.Truncate(time.Millisecond))
	}

	msg.Ack(false)
}
