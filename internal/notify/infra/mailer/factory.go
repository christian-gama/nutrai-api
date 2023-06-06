package mailer

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/log"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
)

func MakeMailer() mailer.Mailer {
	factory, ok := mailers[env.Mailer.Provider]
	if !ok {
		log.Panic(
			errors.New("mailer provider was not found - please check your environment variables"),
		)
	}

	return factory()
}
