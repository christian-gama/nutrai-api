package consumer

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// SaveException is the handler for the SaveException command.
type SaveExceptionHandler interface {
	Handle()
	ConsumerHandler(e exception.Exception) error
}

// saveExceptionHandlerImpl is the implementation of the SaveExceptionHandler.
type saveExceptionHandlerImpl struct {
	message.Consumer[exception.Exception]
	repo.Exception
}

// NewSaveException creates a new SaveException.
func NewSaveExceptionHandler(
	consumer message.Consumer[exception.Exception],
	exceptionRepo repo.Exception,
) SaveExceptionHandler {
	errutil.MustBeNotEmpty("message.Consumer", consumer)
	errutil.MustBeNotEmpty("repo.Exception", exceptionRepo)

	return &saveExceptionHandlerImpl{consumer, exceptionRepo}
}

// Handle handles the event.
func (j *saveExceptionHandlerImpl) Handle() {
	j.Consumer.Handle(j.ConsumerHandler)
}

// ConsumerHandler handles the event.
func (j *saveExceptionHandlerImpl) ConsumerHandler(e exception.Exception) error {
	_, err := j.Save(context.Background(), repo.SaveExceptionInput{
		Exception: &e,
	})
	if err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}
