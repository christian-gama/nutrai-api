package command

import (
	"context"

	"github.com/christian-gama/nutrai-api/internal/core/domain/command"
	"github.com/christian-gama/nutrai-api/internal/core/domain/message"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// RecoveryHandler is the handler for the Recovery command.
type RecoveryHandler = command.Handler[*RecoveryInput]

// recoveryHandlerImpl is the implementation of the RecoveryHandler.
type recoveryHandlerImpl struct {
	message.Publisher[RecoveryInput]
}

// NewRecoveryHandler creates a new RecoveryHandler.
func NewRecoveryHandler(
	publisher message.Publisher[RecoveryInput],
) RecoveryHandler {
	errutil.MustBeNotEmpty("message.Publisher", publisher)

	return &recoveryHandlerImpl{publisher}
}

// Handle handles the command.
func (c *recoveryHandlerImpl) Handle(ctx context.Context, input *RecoveryInput) error {
	c.Publisher.Handle(*input)
	return nil
}
