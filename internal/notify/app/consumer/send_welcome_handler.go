package consumer

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// SaveException is the handler for the event.
type SendWelcomeHandler interface {
	Handle()
	ConsumerHandler(u user.User) error
}

// sendWelcomeHandlerImpl is the implementation of the SendWelcomeHandler interface.
type sendWelcomeHandlerImpl struct {
	message.Consumer[user.User]
	mailer.Mailer
}

// NewSaveException creates a new SendWelcomeHandler.
func NewSendWelcomeHandler(
	consumer message.Consumer[user.User],
	mailer mailer.Mailer,
) SendWelcomeHandler {
	errutil.MustBeNotEmpty("message.Consumer", consumer)
	errutil.MustBeNotEmpty("mailer.Mailer", mailer)

	return &sendWelcomeHandlerImpl{consumer, mailer}
}

// Handle handles the event.
func (j *sendWelcomeHandlerImpl) Handle() {
	j.Consumer.Handle(j.ConsumerHandler)
}

// ConsumerHandler handles the event.
func (j *sendWelcomeHandlerImpl) ConsumerHandler(u user.User) error {
	mail, err := mail.NewMail().
		SetContext(value.Context{"Name": u.Name}).
		SetSubject("Welcome to Nutrai!").
		SetTo(&value.To{Email: u.Email.String(), Name: u.Name.String()}).
		SetTemplatePath(mail.WelcomeTemplate).
		Validate()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if err := j.Mailer.Send(context.Background(), mail); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
