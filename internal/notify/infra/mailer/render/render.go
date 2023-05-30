package render

import (
	"bytes"
	"text/template"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/k3a/html2text"
)

// Render struct is used to load and Render templates, and write text to HTML.
type Render struct {
	template *template.Template
	buffer   bytes.Buffer
}

// New creates a new render.
func New() *Render {
	return &Render{
		buffer: bytes.Buffer{},
	}
}

// LoadTemplate reads and parses a template from given path. Panics if the template cannot be
// loaded.
func (r *Render) LoadTemplate(templatePath *value.Template) *Render {
	if templatePath == nil {
		return &Render{}
	}

	template, err := template.ParseFiles(templatePath.Path)
	if err != nil {
		panic(errors.InternalServerError("failed to load '%s' template", templatePath.Path))
	}

	return &Render{
		template: template,
	}
}

// Render renders the template if it exists, otherwise it writes the body.
func (r *Render) Render(context *value.Context) *Render {
	r.buffer.Reset()

	if r.template != nil {
		return r.renderWithTemplate(context)
	}

	return r.writeToBuffer(context)
}

// ToHTML returns the HTML string.
func (r *Render) ToHTML() string {
	return r.buffer.String()
}

// ToPlainText returns the plain text string.
func (r *Render) ToPlainText() string {
	return html2text.HTML2Text(r.buffer.String())
}

// writeToBuffer writes the body.
func (r *Render) writeToBuffer(context *value.Context) *Render {
	_, err := r.buffer.Write([]byte(context.Body))
	if err != nil {
		panic(errors.InternalServerError("failed to write body"))
	}

	return r
}

// renderWithTemplate renders the template.
func (r *Render) renderWithTemplate(context *value.Context) *Render {
	if err := r.template.Execute(&r.buffer, context); err != nil {
		panic(errors.InternalServerError("failed to render template"))
	}

	return r
}
