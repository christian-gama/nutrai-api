package mailer

import (
	"bytes"
	"text/template"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/k3a/html2text"
)

type renderer struct {
	template *template.Template
	writer   bytes.Buffer
}

func loadTemplate(templatePath *value.TemplatePath) *renderer {
	template, err := template.ParseFiles(templatePath.Path())
	if err != nil {
		panic(errors.InternalServerError("failed to load '%s' template", templatePath.Path()))
	}

	return &renderer{
		template: template,
	}
}

func (b *renderer) render(context *value.Context) *renderer {
	var writer bytes.Buffer = bytes.Buffer{}

	if err := b.template.Execute(&writer, map[string]any{
		"Body": context.Body,
	}); err != nil {
		panic(errors.InternalServerError("failed to render template"))
	}

	b.writer = writer

	return b
}

func (b *renderer) toHTML() string {
	return b.writer.String()
}

func (b *renderer) toPlainText() string {
	return html2text.HTML2Text(b.writer.String())
}
