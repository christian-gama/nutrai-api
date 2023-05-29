package mailer

import (
	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"gopkg.in/gomail.v2"
)

type mailtrapMailer struct{}

func newMailtrapMailer() mailer.Mailer {
	return &mailtrapMailer{}
}

// Send implements mailer.Mailer.
func (m *mailtrapMailer) Send(mail *mail.Mail) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", env.Mailer.From)

	mailer.SetHeader(
		"To",
		slice.Map(mail.To, func(to *value.To) string { return to.Email }).Build()...,
	)

	mailer.SetHeader("Subject", mail.Subject)

	renderer := loadTemplate(mail.TemplatePath).render(mail.Context)
	mailer.SetBody("text/html", renderer.toHTML())
	mailer.AddAlternative("text/plain", renderer.toPlainText())

	for _, attachmentURL := range mail.AttachmentURLs {
		mailer.Attach(attachmentURL)
	}

	dialer := gomail.NewDialer(
		env.Mailtrap.Host,
		env.Mailtrap.Port,
		env.Mailtrap.Username,
		env.Mailtrap.Password,
	)

	return dialer.DialAndSend(mailer)
}
