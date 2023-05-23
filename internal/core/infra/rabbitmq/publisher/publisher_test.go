package publisher_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	mocks "github.com/christian-gama/nutrai-api/testutils/mocks/core/domain/logger"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/stretchr/testify/mock"
)

type PublisherSuite struct {
	suite.Suite
	rmq       *rabbitmq.RabbitMQ
	publisher message.Publisher
	log       *mocks.Logger
}

func TestPublisher(t *testing.T) {
	suite.RunIntegrationTest(t, new(PublisherSuite))
}

func (s *PublisherSuite) SetupTest() {
	s.log = mocks.NewLogger(s.T())
	s.log.On("Infof", mock.Anything, mock.Anything)

	s.rmq = rabbitmq.NewConnection(s.log, "test")

	s.publisher = publisher.NewPublisher(
		s.rmq,
		publisher.WithExchange("test"),
		publisher.WithRoutingKey(event.New("test", event.Save)),
	)
}

func (s *PublisherSuite) TestNewPublisher() {
	s.NotNil(s.publisher)
}

func (s *PublisherSuite) TestPublish() {
	msg := []byte("test-message")

	s.NotPanics(func() {
		s.publisher.Handle(msg)
	})
}
