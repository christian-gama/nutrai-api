package mailer

import (
	"context"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer/render"
	sendgrid "github.com/sendgrid/sendgrid-go"
	sendgridmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridMailer struct {
	client *sendgrid.Client
	render *render.Render
}

func newSendgridMailer() mailer.Mailer {
	return &sendgridMailer{
		client: sendgrid.NewSendClient(env.Sendgrid.ApiKey),
		render: render.New(),
	}
}

// Send implements mailer.Mailer.
func (s *sendgridMailer) Send(ctx context.Context, mail *mail.Mail) error {
	_, err := s.client.SendWithContext(ctx, s.message(mail))
	if err != nil {
		return err
	}

	return err
}

func (s *sendgridMailer) message(mail *mail.Mail) *sendgridmail.SGMailV3 {
	from := sendgridmail.NewEmail(env.Mailer.FromName, env.Mailer.From)
	to := sendgridmail.NewEmail(mail.To[0].Name, mail.To[0].Email)

	render := s.render.Render(mail)

	return sendgridmail.NewSingleEmail(
		from,
		mail.Subject,
		to,
		render.ToPlainText(),
		render.ToHTML(),
	)
}
