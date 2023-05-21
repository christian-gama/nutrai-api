package rabbitmq_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/logger"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type RabbitMQSuite struct {
	suite.Suite
}

func TestRabbitMQ(t *testing.T) {
	suite.RunIntegrationTest(t, new(RabbitMQSuite))
}

func (s *RabbitMQSuite) TestNewConnection() {
	log := mocks.NewLogger(s.T())
	log.On("Infof", mock.Anything, mock.Anything)
	rmq := rabbitmq.NewConnection(log, "test")
	defer rmq.Close()

	channel := rmq.ChannelPool()
	defer rmq.ReleaseChannelPool(channel)

	s.NotNil(rmq)
	s.NotNil(channel)
}

func (s *RabbitMQSuite) TestChannelPool() {
	log := mocks.NewLogger(s.T())
	log.On("Infof", mock.Anything, mock.Anything)
	rmq := rabbitmq.NewConnection(log, "test")
	defer rmq.Close()

	channel := rmq.ChannelPool()
	defer rmq.ReleaseChannelPool(channel)

	s.NotNil(channel)
}

func (s *RabbitMQSuite) TestReleaseChannelPool() {
	log := mocks.NewLogger(s.T())
	log.On("Infof", mock.Anything, mock.Anything)
	rmq := rabbitmq.NewConnection(log, "test")
	defer rmq.Close()

	s.NotPanics(func() {
		channel := rmq.ChannelPool()
		rmq.ReleaseChannelPool(channel)
	})
}
