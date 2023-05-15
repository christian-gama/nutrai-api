package hash_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/infra/hash"
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
		hash.MakeHasher()
	})
}