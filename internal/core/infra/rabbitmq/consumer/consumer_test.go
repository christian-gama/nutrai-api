package consumer_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/consumer"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type ConsumerSuite struct {
	suite.Suite
	rmq      *rabbitmq.RabbitMQ
	consumer message.Consumer[any]
}

func TestConsumer(t *testing.T) {
	suite.RunIntegrationTest(t, new(ConsumerSuite))
}

func (s *ConsumerSuite) SetupSuite() {
	s.rmq = rabbitmq.NewConn("test")
	s.consumer = consumer.NewConsumer[any](
		s.rmq,
		consumer.WithExchangeName("test"),
		consumer.WithRoutingKey(event.New("test", event.Save)),
		consumer.WithArgs(nil),
		consumer.WithAutoDelete(false),
		consumer.WithDurable(false),
		consumer.WithInternal(false),
		consumer.WithKind(rabbitmq.ExchangeDirect),
		consumer.WithNoWait(false),
		consumer.WithQueueArgs(nil),
		consumer.WithQueueAutoDelete(false),
		consumer.WithQueueDurable(false),
		consumer.WithQueueExclusive(false),
		consumer.WithQueueName("test"),
		consumer.WithQueueNoWait(false),
	)
}

func (s *ConsumerSuite) TearDownSuite() {
	s.rmq.Close()
}

func (s *ConsumerSuite) TestNewConsumer() {
	s.NotNil(s.consumer)
}
