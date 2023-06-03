package persistence_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type TokenSuite struct {
	suite.Suite
}

func TestTokenSuite(t *testing.T) {
	t.Skip()
	suite.RunIntegrationTest(t, new(TokenSuite))
}

func (s *TokenSuite) TestToken() {}
