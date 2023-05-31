package mailer

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
)

// Mailer is an interface for sending emails.
type Mailer interface {
	// Send sends an email.
	Send(ctx context.Context, mail *mail.Mail) error
}
