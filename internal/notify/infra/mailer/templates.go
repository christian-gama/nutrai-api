package mailer

import (
	"os"
	"path"

	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/internal/notify/module"
)

func Template() *value.Template {
	return value.NewTemplate().
		SetBaseDir(path.Join(os.Getenv("PWD"), "internal", module.Module.Name(), "infra", "mailer", "templates")).
		SetExt(".html")
}

var (
	Welcome       = Template().SetPath("welcome")
	ResetPassword = Template().SetPath("reset_password")
)
