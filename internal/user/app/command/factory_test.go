package command_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/user/app/command"
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
		command.MakeChangePasswordHandler()
		command.MakeCheckCredentialsHandler()
		command.MakeDeleteUserHandler()
		command.MakeSavePatientHandler()
		command.MakeUpdatePatientHandler()
	})
}
