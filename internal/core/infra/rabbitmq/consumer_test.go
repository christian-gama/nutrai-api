package rabbitmq_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	loggerMock "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/logger"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type ConsumerSuite struct {
	suite.Suite
	rmq      *rabbitmq.RabbitMQ
	consumer message.Consumer
	log      *loggerMock.Logger
}

func TestConsumer(t *testing.T) {
	suite.RunIntegrationTest(t, new(ConsumerSuite))
}

func (s *ConsumerSuite) SetupSuite() {
	s.log = loggerMock.NewLogger(s.T())
	s.log.On("Infof", mock.Anything, mock.Anything)
	s.rmq = rabbitmq.NewConnection(s.log, "test")
	s.consumer = rabbitmq.NewConsumer(
		s.rmq,
		"test-exchange",
		"test-routing-key",
		"test-queue",
		s.log,
	)
}

func (s *ConsumerSuite) TearDownSuite() {
	s.rmq.Close()
}

func (s *ConsumerSuite) TestNewConsumer() {
	s.NotNil(s.consumer)
}
