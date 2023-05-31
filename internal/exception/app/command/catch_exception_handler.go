package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// CatchExceptionHandler is the handler for the CatchException command.
type CatchExceptionHandler = command.Handler[*CatchExceptionInput]

// catchExceptionHandlerImpl is the implementation of the CatchExceptionHandler.
type catchExceptionHandlerImpl struct {
	message.Publisher[exception.Exception]
}

// NewCatchExceptionHandler creates a new CatchExceptionHandler.
func NewCatchExceptionHandler(
	publisher message.Publisher[exception.Exception],
) CatchExceptionHandler {
	errutil.MustBeNotEmpty("message.Publisher", publisher)

	return &catchExceptionHandlerImpl{publisher}
}

// Handle handles the command.
func (c *catchExceptionHandlerImpl) Handle(ctx context.Context, input *CatchExceptionInput) error {
	c.Publisher.Handle(*exception.New().SetMessage(input.Message).SetStack(input.Stack))
	return nil
}
