package publisher

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
)

// Separate the connections for publishers and consumers to achieve high throughput. RabbitMQ can
// apply back pressure on the TCP connection when the publisher is sending too many messages for the
// server to handle. If you consume on the same TCP connection, the server might not receive the
// message acknowledgments from the client, thus effecting the consume performance. With a lower
// consume speed, the server will be overwhelmed.
// Reference:
// https://www.cloudamqp.com/blog/part1-rabbitmq-best-practice.html#separate-connections-for-publisher-and-consumer
var (
	publisherConnection *rabbitmq.RabbitMQ
)

func MakePublisher[Data any](opts ...func(*options)) message.Publisher[Data] {
	if publisherConnection == nil {
		publisherConnection = rabbitmq.NewConn("publisher")
	}

	return NewPublisher[Data](publisherConnection, opts...)
}
