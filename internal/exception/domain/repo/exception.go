package repo

import (
	"context"
	"time"

	"github.com/christian-gama/nutrai-api/internal/exception/domain/model/exception"
)

// SaveExceptionInput is the input for the Save method.
type SaveExceptionInput struct {
	Exception *exception.Exception
}

// DeleteOldExceptionInput is the input for the Delete method.
type DeleteOldExceptionInput struct {
	BeforeDate time.Time
}

// Exception is the interface that wraps the basic Exception methods.
type Exception interface {
	// DeleteOld deletes all exceptions older than the given date.
	DeleteOld(ctx context.Context, input DeleteOldExceptionInput) error

	// Save saves the given exception.
	Save(ctx context.Context, input SaveExceptionInput) (*exception.Exception, error)
}
