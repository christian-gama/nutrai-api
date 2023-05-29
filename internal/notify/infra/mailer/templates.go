package mailer

import (
	"os"
	"path"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/module"
)

var templateDir = path.Join(
	os.Getenv("PWD"),
	"internal",
	module.Module.Name(),
	"infra",
	"mailer",
	"templates",
)

var Welcome *value.TemplatePath = value.NewTemplatePath(templateDir, "welcome.tmpl")
