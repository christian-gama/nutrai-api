package rabbitmq

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/config/env"
)

func uri() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/",
		env.RabbitMQ.User,
		env.RabbitMQ.Password,
		env.RabbitMQ.Host,
		env.RabbitMQ.Port,
	)
}
