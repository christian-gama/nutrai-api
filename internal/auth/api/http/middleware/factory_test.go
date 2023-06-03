package middleware_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/middleware"
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
		middleware.MakeAuth()
	})
}
