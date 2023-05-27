package command_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
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
		command.MakeLogoutHandler()
		command.MakeChangePasswordHandler()
		command.MakeDeleteMeHandler()
		command.MakeSaveUserHandler()
		command.MakeCheckCredentialsHandler()
	})
}
