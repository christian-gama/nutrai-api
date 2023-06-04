package command_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/internal/diet/app/command"
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
		command.MakeSaveDietHandler()
		command.MakeSavePlanHandler()
	})
}
