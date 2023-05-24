package consumer

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/repo"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// SaveException is the handler for the SaveException command.
type SaveExceptionHandler interface {
	Handle()
	ConsumerHandler(body []byte) error
}

// saveExceptionHandlerImpl is the implementation of the SaveExceptionHandler.
type saveExceptionHandlerImpl struct {
	message.Consumer
	repo.Exception
}

// NewSaveException creates a new SaveException.
func NewSaveExceptionHandler(
	consumer message.Consumer,
	exceptionRepo repo.Exception,
) SaveExceptionHandler {
	if consumer == nil {
		panic(errors.New("message.Consumer is required"))
	}

	if exceptionRepo == nil {
		panic(errors.New("repo.Exception is required"))
	}

	return &saveExceptionHandlerImpl{consumer, exceptionRepo}
}

// Handle handles the event.
func (j *saveExceptionHandlerImpl) Handle() {
	j.Consumer.Handle(j.ConsumerHandler)
}

// ConsumerHandler handles the event.
func (j *saveExceptionHandlerImpl) ConsumerHandler(body []byte) error {
	var e exception.Exception
	if err := json.Unmarshal(body, &e); err != nil {
		return errutil.InternalServerError(err.Error())
	}

	_, err := j.Save(context.Background(), repo.SaveExceptionInput{
		Exception: &e,
	})
	if err != nil {
		return errutil.InternalServerError(err.Error())
	}

	return nil
}
