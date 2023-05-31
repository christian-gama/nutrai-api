package publisher_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/domain/event"
	"github.com/christian-gama/nutrai-api/internal/core/infra/rabbitmq/publisher"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type FactorySuite struct {
	suite.Suite
}

func TestFactorySuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(FactorySuite))
}

func (s *FactorySuite) TestFactory() {
	s.NotPanics(func() {
		publisher.MakePublisher[any](
			publisher.WithExchangeName("test"),
			publisher.WithRoutingKey(event.New[any]("test", event.Save)),
		)
	})
}
