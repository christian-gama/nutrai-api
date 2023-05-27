package mailer_test

import (
	"testing"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer"
	"github.com/christian-gama/nutrai-api/testutils/envutil"
	"github.com/christian-gama/nutrai-api/testutils/suite"
)

type FactorySuite struct {
	suite.Suite
}

func TestFactorySuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(FactorySuite))
}

func (s *FactorySuite) TestFactory() {
	s.Run("should not panic if success", func() {
		s.NotPanics(func() {
			mailer.MakeMailer()
		})
	})

	s.Run("should panic if mailer provider is invalid", func() {
		s.Panics(func() {
			reset := envutil.Reset()
			defer reset()
			env.Mailer.Provider = env.MailerProvider("invalid")
			mailer.MakeMailer()
		})
	})
}
