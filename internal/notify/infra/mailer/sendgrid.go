package mailer

import (
	"context"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	sendgrid "github.com/sendgrid/sendgrid-go"
	sendgridmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridMailer struct{}

// Send implements mailer.Mailer.
func (s *sendgridMailer) Send(mail *mail.Mail) error {
	from := sendgridmail.NewEmail(env.Mailer.FromName, env.Mailer.From)

	subject := mail.Subject

	to := sendgridmail.NewEmail(mail.To[0].Name, mail.To[0].Email)

	renderer := loadTemplate(mail.TemplatePath).render(mail.Context)
	message := sendgridmail.NewSingleEmail(
		from,
		subject,
		to,
		renderer.toPlainText(),
		renderer.toHTML(),
	)

	client := sendgrid.NewSendClient(env.Sendgrid.ApiKey)

	_, err := client.SendWithContext(context.Background(), message)
	if err != nil {
		return err
	}

	return err
}

func newSendgridMailer() mailer.Mailer {
	return &sendgridMailer{}
}
