package middleware

import (
	_errors "errors"
	"runtime/debug"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/gin-gonic/gin"
)

// RecoveryHandlerMiddleware is a Gin middleware that recovers from any panics
// and writes a 500 error if there was one. It also handles exceptions and saves
// them using the provided exception handler.
type Recovery = middleware.Middleware

// NewRecovery is a constructor function that returns a RecoveryHandler that wraps around a provided
// recoveryHandler. The function panics if a nil recoveryHandler is provided. It is used to create
// an instance of middleware that handles and logs exceptions during request handling.
func NewRecovery(
	recoveryHandler command.RecoveryHandler,
) Recovery {
	errutil.MustBeNotEmpty("command.RecoveryHandler", recoveryHandler)

	m := &recoveryHandlerImpl{
		recoveryHandler: recoveryHandler,
	}

	return middleware.NewMiddleware(m.Handle)
}

// recoveryHandlerImpl is an implementation of the RecoveryHandler middleware.
// It delegates the exception handling to a provided recoveryHandler.
type recoveryHandlerImpl struct {
	recoveryHandler command.RecoveryHandler
}

// Handle is a middleware function that recovers from panics that occur during request handling,
// and passes the error to the recoveryHandler for logging.
func (m *recoveryHandlerImpl) Handle(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			m.handleException(ctx, r)
		}
	}()

	ctx.Next()
}

// handleException handles recovered panics. It logs the exception using recoveryHandler and
// sends a 500 Internal Server Error response to the client.
func (m *recoveryHandlerImpl) handleException(ctx *gin.Context, r any) {
	message := m.getErrorMessage(r)

	m.recoveryHandler.Handle(ctx.Request.Context(),
		&command.RecoveryInput{
			Message: message,
			Stack:   string(debug.Stack()),
		},
	)

	if env.IsProduction {
		message = "something went wrong, please try again later"
	}

	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		response.Error(errors.InternalServerError(message)),
	)
}

// getErrorMessage constructs the error message from the recovered panic. It uses type assertion to
// check if the panic was caused by an error, and if not, a generic error message is used.
func (m *recoveryHandlerImpl) getErrorMessage(r any) string {
	switch err := r.(type) {
	case error:
		return m.getErrorFromException(err).Error()

	default:
		return errors.InternalServerError("something went wrong, please try again later").Error()
	}
}

// getErrorFromException constructs an error message from an exception. It checks if the
// error is of type errors.ErrInternal and if not, wraps the error into an errors.ErrInternal.
func (m *recoveryHandlerImpl) getErrorFromException(err error) error {
	var errInternal *errors.ErrInternalServerError
	if !_errors.As(err, &errInternal) {
		return errors.InternalServerError(err.Error())
	}

	return err
}
