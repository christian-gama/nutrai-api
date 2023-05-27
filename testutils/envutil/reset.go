package envutil

import "github.com/christian-gama/nutrai-api/config/env"

func Reset() func() {
	originalAppEnv := *env.App
	originalDBEnv := *env.DB
	originalConfigEnv := *env.Config
	originalJWTEnv := *env.Jwt
	originalMailerEnv := *env.Mailer
	originalSendgridEnv := *env.Sendgrid
	originalMailtrapEnv := *env.Mailtrap

	return func() {
		env.App = &originalAppEnv
		env.DB = &originalDBEnv
		env.Config = &originalConfigEnv
		env.Jwt = &originalJWTEnv
		env.Mailer = &originalMailerEnv
		env.Sendgrid = &originalSendgridEnv
		env.Mailtrap = &originalMailtrapEnv
	}
}
