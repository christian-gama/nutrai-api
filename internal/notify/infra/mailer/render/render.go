package render

import (
	"bytes"
	"path"
	"text/template"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/k3a/html2text"
)

// Render struct is used to load and Render templates, and write text to HTML.
type Render struct {
	template *template.Template
	buffer   bytes.Buffer
	base     string
}

// New creates a new render.
func New() *Render {
	return &Render{
		buffer: bytes.Buffer{},
	}
}

// LoadTemplate reads and parses a template from given path. Panics if the template cannot be
// loaded.
func (r *Render) loadTemplate(templatePath *value.Template) {
	if templatePath == nil {
		return
	}

	r.base = path.Join(env.Mailer.TemplatePath, "__base__.html")

	template, err := template.ParseFiles(r.base, templatePath.Path())
	if err != nil {
		panic(errors.InternalServerError("failed to load '%s' template", templatePath.Path()))
	}

	r.template = template
}

// Render renders the template if it exists, otherwise it writes the body.
func (r *Render) Render(mail *mail.Mail) *rendered {
	r.loadTemplate(mail.Template)
	r.buffer.Reset()
	return r.renderWithTemplate(mail.Context)
}

// renderWithTemplate renders the template.
func (r *Render) renderWithTemplate(context value.Context) *rendered {
	if err := r.template.ExecuteTemplate(&r.buffer, "base", context); err != nil {
		panic(errors.InternalServerError("failed to render template: %v", err))
	}

	return &rendered{buffer: r.buffer}
}

type rendered struct {
	buffer bytes.Buffer
}

// ToHTML returns the HTML string.
func (r *rendered) ToHTML() string {
	return r.buffer.String()
}

// ToPlainText returns the plain text string.
func (r *rendered) ToPlainText() string {
	return html2text.HTML2Text(r.buffer.String())
}
