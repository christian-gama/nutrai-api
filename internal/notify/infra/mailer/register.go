package mailer

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
)

var mailers = map[env.MailerProvider]func() mailer.Mailer{}

func register(provider env.MailerProvider, factory func() mailer.Mailer) {
	mailers[provider] = factory
}

func init() {
	register(env.MailerProviderMailtrap, newMailtrapMailer)
	register(env.MailerProviderSendgrid, newSendgridMailer)
}
