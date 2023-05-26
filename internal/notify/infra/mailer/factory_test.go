package mailer_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
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
		mailer.MakeMailer()
	})

	s.Panics(func() {
		original := env.Mailer.Provider
		env.Mailer.Provider = env.MailerProvider("invalid")
		mailer.MakeMailer()
		env.Mailer.Provider = original
	})
}
