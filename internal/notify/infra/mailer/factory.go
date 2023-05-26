package mailer

import (
	"errors"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
)

func MakeMailer() mailer.Mailer {
	factory, ok := mailers[env.Mailer.Provider]
	if !ok {
		panic(errors.New("mailer provider not found"))
	}

	return factory()
}
