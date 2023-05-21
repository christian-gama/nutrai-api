package rabbitmq

import (
	"fmt"
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
	"github.com/christian-gama/nutrai-api/internal/core/infra/env"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ is a wrapper for amqp.Connection.
type RabbitMQ struct {
	conn *amqp.Connection
	ch   chan *amqp.Channel
}

// NewConnection creates a new RabbitMQ connection.
func NewConnection(log logger.Logger, mode string) *RabbitMQ {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		env.RabbitMQ.User,
		env.RabbitMQ.Password,
		env.RabbitMQ.Host,
		env.RabbitMQ.Port,
	)

	log.Infof("Connecting to RabbitMQ (%s)", mode)
	maxAttempts := 3
	var err error
	var conn *amqp.Connection
	for attempts := 0; attempts < maxAttempts; attempts++ {
		conn, err = amqp.Dial(uri)
		if err == nil {
			ch := makeChannelPool(conn)
			return &RabbitMQ{conn, ch}
		}

		time.Sleep(1 * time.Second)
	}

	log.Panic(fmt.Errorf("could not connect to RabbitMQ: %w", err))
	return nil
}

// Close closes the RabbitMQ connection.
func (r *RabbitMQ) Close() {
	r.conn.Close()
}

// ChannelPool returns a channel from the pool.
func (r *RabbitMQ) ChannelPool() *amqp.Channel {
	return <-r.ch
}

// ReleaseChannelPool releases a channel to the pool.
func (r *RabbitMQ) ReleaseChannelPool(ch *amqp.Channel) {
	r.ch <- ch
}

// makeChannelPool creates a channel pool.
func makeChannelPool(conn *amqp.Connection) chan *amqp.Channel {
	ch := make(chan *amqp.Channel, 1)

	select {
	case channel := <-ch:
		if err := channel.Confirm(false); err != nil {
			panic(err)
		}
		ch <- channel

	default:
		channel, err := conn.Channel()
		if err != nil {
			panic(err)
		}

		if err := channel.Confirm(false); err != nil {
			panic(err)
		}
		ch <- channel
	}

	return ch
}
