package mailer

import (
	"context"
	"os"

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

	v3 := sendgridmail.NewV3MailInit(
		from,
		mail.Subject,
		to,
		sendgridmail.NewContent("text/plain", render.ToPlainText()),
		sendgridmail.NewContent("text/html", render.ToHTML()),
	)

	for _, attachment := range mail.Attachments {
		a := sendgridmail.NewAttachment()
		a.SetFilename(attachment.Filename)
		a.SetContentID(attachment.ContentID())
		a.SetDisposition(attachment.Disposition)
		a.SetType(attachment.ContentType())
		a.SetContent(attachment.Content(os.ReadFile))
		v3.AddAttachment(a)
	}

	return v3
}
