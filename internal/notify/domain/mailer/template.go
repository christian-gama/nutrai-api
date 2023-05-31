package mailer

import value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"

var (
	Welcome       = value.NewTemplate().SetPath("welcome")
	ResetPassword = value.NewTemplate().SetPath("reset_password")
)
