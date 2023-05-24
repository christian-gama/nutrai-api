package mail

import (
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Mail represents an email message.
type Mail struct {
	To       []value.Email  `json:"to" faker:"-"`
	From     value.Email    `json:"from" faker:"email"`
	Subject  value.Subject  `json:"subject" faker:"len=50"`
	Body     value.Body     `json:"body" faker:"len=100"`
	Template value.Template `json:"template" faker:"len=100"`
}

// NewMail creates a new Mail.
func NewMail() *Mail {
	return &Mail{}
}

// Validate validates the Mail fields.
func (m *Mail) Validate() (*Mail, error) {
	var errs *errutil.Error

	if m.Body == "" {
		errs = errutil.Append(errs, errutil.Required("Body"))
	}

	if len(m.To) == 0 {
		errs = errutil.Append(errs, errutil.Required("To"))
	}

	for _, to := range m.To {
		if to == "" {
			errs = errutil.Append(errs, errutil.Required("To"))
		}
	}

	if m.From == "" {
		errs = errutil.Append(errs, errutil.Required("From"))
	}

	if m.Subject == "" {
		errs = errutil.Append(errs, errutil.Required("Subject"))
	}

	if m.Template == "" {
		errs = errutil.Append(errs, errutil.Required("Template"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}

// SetTo sets the To field.
func (m *Mail) SetTo(to []value.Email) *Mail {
	m.To = to
	return m
}

// SetFrom sets the From field.
func (m *Mail) SetFrom(from value.Email) *Mail {
	m.From = from
	return m
}

// SetSubject sets the Subject field.
func (m *Mail) SetSubject(subject value.Subject) *Mail {
	m.Subject = subject
	return m
}

// SetBody sets the Body field.
func (m *Mail) SetBody(body value.Body) *Mail {
	m.Body = body
	return m
}

// SetTemplate sets the Template field.
func (m *Mail) SetTemplate(template value.Template) *Mail {
	m.Template = template
	return m
}
