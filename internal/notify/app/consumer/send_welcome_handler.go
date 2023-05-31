package consumer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	m "github.com/christian-gama/nutrai-api/internal/notify/domain/mailer"
	"github.com/christian-gama/nutrai-api/internal/notify/domain/model/mail"
	value "github.com/christian-gama/nutrai-api/internal/notify/domain/value/mail"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// SaveException is the handler for the event.
type SendWelcomeHandler interface {
	Handle()
	ConsumerHandler(body []byte) error
}

// sendWelcomeHandlerImpl is the implementation of the SendWelcomeHandler interface.
type sendWelcomeHandlerImpl struct {
	message.Consumer
	m.Mailer
}

// NewSaveException creates a new SendWelcomeHandler.
func NewSendWelcomeHandler(
	consumer message.Consumer,
	mailer m.Mailer,
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
func (j *sendWelcomeHandlerImpl) ConsumerHandler(body []byte) error {
	var user user.User
	if err := json.Unmarshal(body, &user); err != nil {
		return errors.InternalServerError(err.Error())
	}

	mail := mail.NewMail().
		SetContext(value.NewContext().
			SetBody(fmt.Sprintf("Dear %s,\n\nCongratulations on taking a bold step towards your health and wellness journey with us. We are excited to have you on board!\n\nOur AI-powered application is designed to help you create personalized diet plans based on your unique health and lifestyle needs. The program tailors every meal plan to align with your individual goals, whether it's weight loss, muscle gain, or maintaining a balanced lifestyle.\n\nWith our tool, you no longer have to worry about navigating the complex world of nutrition alone. Our AI will help you create a balanced diet that not only suits your taste preferences but also provides the nutrition your body needs.\n\nTo get started, simply input your current dietary habits, lifestyle, and health goals, and let the AI do the rest! It will generate a daily, weekly, or monthly diet plan for you, keeping in mind any dietary restrictions you may have.", user.Name)),
		).
		SetSubject("Welcome to Nutrai!").
		SetTo([]*value.To{
			{Email: user.Email.String(), Name: user.Name.String()},
		}).
		SetTemplatePath(mailer.Welcome)

	if err := j.Mailer.Send(context.Background(), mail); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
