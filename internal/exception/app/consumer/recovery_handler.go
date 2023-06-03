package consumer

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Recovery is the handler for the Recovery command.
type RecoveryHandler interface {
	Handle()
	ConsumerHandler(input command.RecoveryInput) error
}

// saveExceptionHandlerImpl is the implementation of the RecoveryHandler.
type saveExceptionHandlerImpl struct {
	message.Consumer[command.RecoveryInput]
	repo.Exception
}

// NewRecovery creates a new Recovery.
func NewRecoveryHandler(
	consumer message.Consumer[command.RecoveryInput],
	exceptionRepo repo.Exception,
) RecoveryHandler {
	errutil.MustBeNotEmpty("message.Consumer", consumer)
	errutil.MustBeNotEmpty("repo.Exception", exceptionRepo)

	return &saveExceptionHandlerImpl{consumer, exceptionRepo}
}

// Handle handles the event.
func (j *saveExceptionHandlerImpl) Handle() {
	j.Consumer.Handle(j.ConsumerHandler)
}

// ConsumerHandler handles the event.
func (j *saveExceptionHandlerImpl) ConsumerHandler(input command.RecoveryInput) error {
	_, err := j.Save(context.Background(), repo.SaveExceptionInput{
		Exception: exception.New().SetMessage(input.Message).SetStack(input.Stack),
	})
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
