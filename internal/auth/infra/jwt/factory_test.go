package jwt_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/infra/jwt"
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
		jwt.MakeAccessTokenGenerator()
		jwt.MakeAccessVerifier()
		jwt.MakeRefreshTokenGenerator()
		jwt.MakeRefreshVerifier()
	})
}
