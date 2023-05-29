package mail

import (
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Mail represents an email message.
type Mail struct {
	To             []*value.To         `json:"to" faker:"-"`
	Subject        string              `json:"subject" faker:"len=50"`
	Context        *value.Context      `json:"context" faker:"-"`
	TemplatePath   *value.TemplatePath `json:"templatePath" faker:"-"`
	AttachmentURLs []string            `json:"attachmentURLs" faker:"-"`
}

// NewMail creates a new Mail.
func NewMail() *Mail {
	return &Mail{}
}

// Validate validates the Mail fields.
func (m *Mail) Validate() (*Mail, error) {
	var errs *errutil.Error

	if m.Context == nil {
		errs = errutil.Append(errs, errors.Required("Context"))
	}

	if len(m.To) == 0 {
		errs = errutil.Append(errs, errors.Required("To"))
	}

	for _, to := range m.To {
		if to.Email == "" {
			errs = errutil.Append(errs, errors.Required("To.Email"))
		}

		if to.Name == "" {
			errs = errutil.Append(errs, errors.Required("To.Name"))
		}
	}

	if m.Subject == "" {
		errs = errutil.Append(errs, errors.Required("Subject"))
	}

	if m.TemplatePath == nil {
		errs = errutil.Append(errs, errors.Required("TemplatePath"))
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

// SetContext sets the Context field.
func (m *Mail) SetContext(body, head string) *Mail {
	m.Context = value.NewContext(body, head)
	return m
}

// SetTemplatePath sets the Template field.
func (m *Mail) SetTemplatePath(paths *value.TemplatePath) *Mail {
	m.TemplatePath = paths
	return m
}
