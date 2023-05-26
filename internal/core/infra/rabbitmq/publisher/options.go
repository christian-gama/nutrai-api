package publisher

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	amqp "github.com/rabbitmq/amqp091-go"
)

type options struct {
	ExchangeName string
	RoutingKey   string
	Kind         rabbitmq.Exchange
	Durable      bool
	AutoDelete   bool
	Internal     bool
	NoWait       bool
	Args         amqp.Table
	Mandatory    bool
	Immediate    bool
	ContentType  rabbitmq.ContentType
}

func WithExchange(exchange rabbitmq.Exchange) func(*options) {
	errutil.MustBeNotEmpty("exchange", exchange)

	return func(o *options) {
		o.ExchangeName = exchange
	}
}

func WithRoutingKey(routingKey event.Event) func(*options) {
	errutil.MustBeNotEmpty("routingKey", routingKey)

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

func WithMandatory(mandatory bool) func(*options) {
	return func(o *options) {
		o.Mandatory = mandatory
	}
}

func WithImmediate(immediate bool) func(*options) {
	return func(o *options) {
		o.Immediate = immediate
	}
}

func WithContentType(contentType rabbitmq.ContentType) func(*options) {
	errutil.MustBeNotEmpty("contentType", contentType)

	return func(o *options) {
		o.ContentType = contentType
	}
}
