package rabbitmq_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type RabbitMQSuite struct {
	suite.Suite
}

func TestRabbitMQ(t *testing.T) {
	suite.RunIntegrationTest(t, new(RabbitMQSuite))
}

func (s *RabbitMQSuite) TestNewConnection() {
	rmq := rabbitmq.NewConn("test")
	defer rmq.Close()

	channel := rmq.ChannelPool()
	defer rmq.ReleaseChannelPool(channel)

	s.NotNil(rmq)
	s.NotNil(channel)
}

func (s *RabbitMQSuite) TestChannelPool() {
	rmq := rabbitmq.NewConn("test")
	defer rmq.Close()

	channel := rmq.ChannelPool()
	defer rmq.ReleaseChannelPool(channel)

	s.NotNil(channel)
}

func (s *RabbitMQSuite) TestReleaseChannelPool() {
	rmq := rabbitmq.NewConn("test")
	defer rmq.Close()

	s.NotPanics(func() {
		channel := rmq.ChannelPool()
		rmq.ReleaseChannelPool(channel)
	})
}
