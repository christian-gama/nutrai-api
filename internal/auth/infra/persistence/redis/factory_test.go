package persistence_test

import (
	"testing"

	persistence "github.com/christian-gama/nutrai-api/internal/auth/infra/persistence/redis"
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
		persistence.MakeRedisToken()
	})
}
