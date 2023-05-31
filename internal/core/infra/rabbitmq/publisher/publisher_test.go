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
	publisher message.Publisher[Data]
	log       *mocks.Logger
}

func TestPublisher(t *testing.T) {
	suite.RunIntegrationTest(t, new(PublisherSuite))
}

func (s *PublisherSuite) SetupTest() {
	s.log = mocks.NewLogger(s.T())
	s.log.On("Loading", mock.Anything, mock.Anything)

	s.rmq = rabbitmq.NewConn(s.log, "test").Open()

	s.publisher = publisher.NewPublisher[Data](
		s.rmq,
		publisher.WithExchangeName("test"),
		publisher.WithRoutingKey(Event),
		publisher.WithArgs(nil),
		publisher.WithAutoDelete(false),
		publisher.WithDurable(false),
		publisher.WithInternal(false),
		publisher.WithKind(rabbitmq.ExchangeDirect),
		publisher.WithNoWait(false),
		publisher.WithMandatory(false),
		publisher.WithImmediate(false),
		publisher.WithContentType(rabbitmq.ContentTypeJSON),
	)
}

func (s *PublisherSuite) TestNewPublisher() {
	s.NotNil(s.publisher)
}

func (s *PublisherSuite) TestPublish() {
	s.NotPanics(func() {
		s.publisher.Handle(Data{Name: "test"})
	})
}

type Data struct {
	Name string `json:"name"`
}

var Event = event.New[Data]("test", event.Save)
