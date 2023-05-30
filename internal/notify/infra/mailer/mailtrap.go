package mailer

import (
	"context"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/infra/mailer/render"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"gopkg.in/gomail.v2"
)

type mailtrapMailer struct {
	msg    *gomail.Message
	dialer *gomail.Dialer
	render *render.Render
}

func newMailtrapMailer() mailer.Mailer {
	return &mailtrapMailer{
		msg: gomail.NewMessage(),
		dialer: gomail.NewDialer(
			env.Mailtrap.Host,
			env.Mailtrap.Port,
			env.Mailtrap.Username,
			env.Mailtrap.Password,
		),
		render: render.New(),
	}
}

// Send implements mailer.Mailer.
func (m *mailtrapMailer) Send(ctx context.Context, mail *mail.Mail) error {
	m.setHeaders(mail)
	m.setBody(mail)
	m.setAttachments(mail)
	return m.dialer.DialAndSend(m.msg)
}

func (m *mailtrapMailer) setHeaders(mail *mail.Mail) {
	m.msg.SetHeader("From", env.Mailer.From)
	m.msg.SetHeader(
		"To",
		slice.Map(mail.To, func(to *value.To) string { return to.Email }).Build()...,
	)
	m.msg.SetHeader("Subject", mail.Subject)
}

func (m *mailtrapMailer) setBody(mail *mail.Mail) {
	render := m.render.LoadTemplate(mail.Template).Render(mail.Context)

	m.msg.SetBody("text/html", render.ToHTML())
	m.msg.AddAlternative("text/plain", render.ToPlainText())
}

func (m *mailtrapMailer) setAttachments(mail *mail.Mail) {
	for _, attachmentURL := range mail.AttachmentURLs {
		m.msg.Attach(attachmentURL)
	}
}
