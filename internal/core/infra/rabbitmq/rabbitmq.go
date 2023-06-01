package rabbitmq

import (
	"time"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"
	"github.com/christian-gama/nutrai-api/pkg/retry"
	amqp "github.com/rabbitmq/amqp091-go"
)

// NewConn creates a new RabbitMQ connection.
func NewConn(mode string) *RabbitMQ {
	log.Loading("Connecting to RabbitMQ (%s)", mode)

	const attempts = 90
	var err error
	rabbitMQ := &RabbitMQ{mode: mode}
	retry.Retry(attempts, time.Second, func() error {
		conn, err := amqp.Dial(uri())
		if err == nil {
			rabbitMQ.conn = conn
			rabbitMQ.ch = makeChannelPool(conn)
		}
		return err
	})

	if err != nil {
		log.Fatalf("\tFailed to connect to RabbitMQ after %d retries: %v", attempts, err)
	}

	return rabbitMQ
}

// RabbitMQ is a wrapper for amqp.Connection.
type RabbitMQ struct {
	conn *amqp.Connection
	ch   chan *amqp.Channel
	mode string
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

// Connection returns the RabbitMQ connection.
func (r *RabbitMQ) Connection() *amqp.Connection {
	return r.conn
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
