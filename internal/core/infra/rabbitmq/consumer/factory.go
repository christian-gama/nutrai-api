package consumer

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
)

// Separate the instances for publishers and consumers to achieve high throughput. RabbitMQ can
// apply back pressure on the TCP connection when the publisher is sending too many messages for the
// server to handle. If you consume on the same TCP connection, the server might not receive the
// message acknowledgments from the client, thus effecting the consume performance. With a lower
// consume speed, the server will be overwhelmed.
// Reference:
// https://www.cloudamqp.com/blog/part1-rabbitmq-best-practice.html#separate-connections-for-publisher-and-consumer
var rabbitmqInstance *rabbitmq.RabbitMQ

func MakeConsumer[Data any](opts ...func(*options)) message.Consumer[Data] {
	if rabbitmqInstance == nil {
		rabbitmqInstance = rabbitmq.NewConn("consumer")
	}

	return NewConsumer[Data](rabbitmqInstance, opts...)
}
