package consumer

import (
	"context"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/app/command"
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
	ConsumerHandler(input command.SaveUserInput) error
}

// sendWelcomeHandlerImpl is the implementation of the SendWelcomeHandler interface.
type sendWelcomeHandlerImpl struct {
	message.Consumer[command.SaveUserInput]
	mailer.Mailer
}

// NewSaveException creates a new SendWelcomeHandler.
func NewSendWelcomeHandler(
	consumer message.Consumer[command.SaveUserInput],
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
func (j *sendWelcomeHandlerImpl) ConsumerHandler(input command.SaveUserInput) error {
	mail, err := mail.NewMail().
		SetContext(value.Context{"Name": input.Name.String(), "Title": fmt.Sprintf("Welcome, %s!", input.Name.String())}).
		SetSubject("Welcome to Nutrai!").
		SetTo(&value.To{Email: input.Email.String(), Name: input.Name.String()}).
		SetAttachments(value.NewAttachment().
			SetFilename(mail.BuildAssetURL("welcome.png")).
			SetDisposition("inline"),
		).
		SetTemplate("welcome").
		Validate()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	if err := j.Mailer.Send(context.Background(), mail); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
