package mail

import (
	"os"
	"path"

	"github.com/christian-gama/nutrai-api/config/env"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Mail represents an email message.
type Mail struct {
	To          []*value.To         `json:"to" faker:"-"`
	Subject     string              `json:"subject" faker:"len=50"`
	Context     value.Context       `json:"context" faker:"-"`
	Template    *value.Template     `json:"templatePath" faker:"-"`
	Attachments []*value.Attachment `json:"attachments" faker:"-"`
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

	if m.Template == nil {
		errs = errutil.Append(errs, errors.Required("Template"))
	}

	if errs.HasErrors() {
		return nil, errs
	}

	return m, nil
}

// SetTo sets the To field.
func (m *Mail) SetTo(to ...*value.To) *Mail {
	m.To = to
	return m
}

// SetSubject sets the Subject field.
func (m *Mail) SetSubject(subject string) *Mail {
	m.Subject = subject
	return m
}

// SetContext sets the Context field.
func (m *Mail) SetContext(context value.Context) *Mail {
	m.Context = context
	return m
}

// SetTemplate sets the Template field. The name should be the base name of the file, such as
// "welcome.html".
func (m *Mail) SetTemplate(name string) *Mail {
	m.Template = value.NewTemplate(name)
	return m
}

// SetAttachmentURLs sets the AttachmentURLs field. The names should be the base name of the file,
// such as "welcome.png".
func (m *Mail) SetAttachments(att ...*value.Attachment) *Mail {
	m.Attachments = att
	return m
}

// BuildAssetURL builds the URL path to the asset.
func BuildAssetURL(name string) string {
	return path.Join(os.Getenv("PWD"), env.Mailer.AssetsPath, name)
}
