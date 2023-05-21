package rabbitmq

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/bench"
	amqp "github.com/rabbitmq/amqp091-go"
)

// consumerImpl is a RabbitMQ consumer implementation.
type consumerImpl struct {
	*RabbitMQ
	exchange   string
	queueName  string
	routingKey string
	log        logger.Logger
}

// NewConsumer creates a new RabbitMQ consumer.
func NewConsumer(
	conn *RabbitMQ,
	exchange string,
	routingKey string,
	queueName string,
	log logger.Logger,
) message.Consumer {
	return &consumerImpl{conn, exchange, queueName, routingKey, log}
}

// Handle handles a message.
func (c *consumerImpl) Handle(handler message.MessageHandler) {
	ch, err := c.conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		c.exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	q, err := ch.QueueDeclare(
		c.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	if err := ch.Qos(1, 0, false); err != nil {
		return
	}

	err = ch.QueueBind(
		q.Name,
		c.routingKey,
		c.exchange,
		false,
		nil,
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

func (c *consumerImpl) handle(msg amqp.Delivery, handler message.MessageHandler) {
	defer func() {
		if r := recover(); r != nil {
			c.log.Warnf("Consumer | %s | Recovered from panic: %s", c.queueName, r)
			msg.Nack(false, false)
		}
	}()

	var err error
	duration := bench.Duration(func() {
		err = handler(msg.Body)
	})

	if err != nil {
		c.log.Warnf("Consumer | %s | Could not process message: %s", c.queueName, err.Error())
	} else {
		c.log.Infof("Consumer | %s | Message processed successfully in %dms", c.queueName, duration.Milliseconds())
	}

	msg.Ack(false)
}
