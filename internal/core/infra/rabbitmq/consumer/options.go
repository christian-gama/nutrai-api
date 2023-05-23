package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type options struct {
	ExchangeName    string
	RoutingKey      string
	Kind            rabbitmq.Exchange
	Durable         bool
	AutoDelete      bool
	Internal        bool
	NoWait          bool
	Args            amqp.Table
	QueueName       string
	QueueDurable    bool
	QueueAutoDelete bool
	QueueExclusive  bool
	QueueNoWait     bool
	QueueArgs       amqp.Table
}

func WithExchange(exchange rabbitmq.Exchange) func(*options) {
	return func(o *options) {
		o.ExchangeName = exchange
	}
}

func WithRoutingKey(routingKey event.Event) func(*options) {
	return func(o *options) {
		o.RoutingKey = routingKey.String()
	}
}

func WithKind(kind string) func(*options) {
	return func(o *options) {
		o.Kind = kind
	}
}

func WithDurable(durable bool) func(*options) {
	return func(o *options) {
		o.Durable = durable
	}
}

func WithAutoDelete(autoDelete bool) func(*options) {
	return func(o *options) {
		o.AutoDelete = autoDelete
	}
}

func WithInternal(internal bool) func(*options) {
	return func(o *options) {
		o.Internal = internal
	}
}

func WithNoWait(noWait bool) func(*options) {
	return func(o *options) {
		o.NoWait = noWait
	}
}

func WithArgs(args amqp.Table) func(*options) {
	return func(o *options) {
		o.Args = args
	}
}

func WithQueueName(queueName string) func(*options) {
	return func(o *options) {
		o.QueueName = queueName
	}
}

func WithQueueDurable(queueDurable bool) func(*options) {
	return func(o *options) {
		o.QueueDurable = queueDurable
	}
}

func WithQueueAutoDelete(queueAutoDelete bool) func(*options) {
	return func(o *options) {
		o.QueueAutoDelete = queueAutoDelete
	}
}

func WithQueueExclusive(queueExclusive bool) func(*options) {
	return func(o *options) {
		o.QueueExclusive = queueExclusive
	}
}

func WithQueueNoWait(queueNoWait bool) func(*options) {
	return func(o *options) {
		o.QueueNoWait = queueNoWait
	}
}

func WithQueueArgs(queueArgs amqp.Table) func(*options) {
	return func(o *options) {
		o.QueueArgs = queueArgs
	}
}