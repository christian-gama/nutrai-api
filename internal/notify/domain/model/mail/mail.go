package mail

import (
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Mail represents an email message.
type Mail struct {
	To             []*value.To `json:"to" faker:"-"`
	Subject        string      `json:"subject" faker:"len=50"`
	PlainText      string      `json:"plainText" faker:"len=100"`
	HTML           string      `json:"html" faker:"len=100"`
	AttachmentURLs []string    `json:"attachmentURLs" faker:"-"`
}

// NewMail creates a new Mail.
func NewMail() *Mail {
	return &Mail{}
}

// Validate validates the Mail fields.
func (m *Mail) Validate() (*Mail, error) {
	var errs *errutil.Error

	if m.PlainText == "" {
		errs = errutil.Append(errs, errutil.Required("Body"))
	}

	if len(m.To) == 0 {
		errs = errutil.Append(errs, errutil.Required("To"))
	}

	for _, to := range m.To {
		if to.Email == "" {
			errs = errutil.Append(errs, errutil.Required("To.Email"))
		}

		if to.Name == "" {
			errs = errutil.Append(errs, errutil.Required("To.Name"))
		}
	}

	if m.Subject == "" {
		errs = errutil.Append(errs, errutil.Required("Subject"))
	}

	if m.HTML == "" {
		errs = errutil.Append(errs, errutil.Required("Template"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}

// SetTo sets the To field.
func (m *Mail) SetTo(to []*value.To) *Mail {
	m.To = to
	return m
}

// SetSubject sets the Subject field.
func (m *Mail) SetSubject(subject string) *Mail {
	m.Subject = subject
	return m
}

// SetPlainText sets the Body field.
func (m *Mail) SetPlainText(body string) *Mail {
	m.PlainText = body
	return m
}

// SetHTML sets the Template field.
func (m *Mail) SetHTML(template string) *Mail {
	m.HTML = template
	return m
}
