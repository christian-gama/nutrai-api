package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/app/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// SaveException is a command handler that saves an exception to the database.
type SaveExceptionHandler = command.Handler[*SaveExceptionInput]

// saveExceptionImpl implements the SaveException interface.
type saveExceptionHandlerImpl struct {
	message.Publisher
}

// NewSaveExceptionHandler creates a new SaveExceptionHandler.
func NewSaveExceptionHandler(publisher message.Publisher) SaveExceptionHandler {
	if publisher == nil {
		panic(errors.New("message.Publisher is required"))
	}

	return &saveExceptionHandlerImpl{publisher}
}

// Handle handles the command.
func (c *saveExceptionHandlerImpl) Handle(ctx context.Context, input *SaveExceptionInput) error {
	encoded, err := json.Marshal(&input)
	if err != nil {
		return errutil.InternalServerError(fmt.Sprintf("failed to encode error: %v", err))
	}

	c.Publisher.Handle(encoded)

	return nil
}
