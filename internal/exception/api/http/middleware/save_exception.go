package middleware

import (
	"errors"
	"runtime/debug"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// SaveExceptionHandlerMiddleware is a Gin middleware that recovers from any panics
// and writes a 500 error if there was one. It also handles exceptions and saves
// them using the provided exception handler.
type SaveException = middleware.Middleware

// NewSaveException is a constructor function that returns a SaveExceptionHandler that wraps
// around a provided saveExceptionHandler. The function panics if a nil saveExceptionHandler is
// provided. It is used to create an instance of middleware that handles and logs exceptions during
// request handling.
func NewSaveException(
	saveExceptionHandler command.SaveExceptionHandler,
) SaveException {
	if saveExceptionHandler == nil {
		panic(errors.New("command.SaveExceptionHandler cannot be nil"))
	}

	m := &saveExceptionHandlerImpl{
		saveExceptionHandler: saveExceptionHandler,
	}

	return middleware.NewMiddleware(m.Handle)
}

// saveExceptionHandlerImpl is an implementation of the SaveExceptionHandler middleware.
// It delegates the exception handling to a provided saveExceptionHandler.
type saveExceptionHandlerImpl struct {
	saveExceptionHandler command.SaveExceptionHandler
}

// Handle is a middleware function that recovers from panics that occur during request handling,
// and passes the error to the saveExceptionHandler for logging.
func (m *saveExceptionHandlerImpl) Handle(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			m.handleException(ctx, r)
		}
	}()

	ctx.Next()
}

// handleException handles recovered panics. It logs the exception using saveExceptionHandler and
// sends a 500 Internal Server Error response to the client.
func (m *saveExceptionHandlerImpl) handleException(ctx *gin.Context, r any) {
	message := m.getErrorMessage(r)

	m.saveExceptionHandler.Handle(ctx.Request.Context(),
		&command.SaveExceptionInput{
			Message: message,
			Stack:   string(debug.Stack()),
		},
	)

	if env.App.Env == env.Production {
		message = "something went wrong, please try again later"
	}

	ctx.AbortWithStatusJSON(
		http.StatusInternalServerError,
		response.Error(errors.New(message)),
	)
}

// getErrorMessage constructs the error message from the recovered panic. It uses type assertion to
// check if the panic was caused by an error, and if not, a generic error message is used.
func (m *saveExceptionHandlerImpl) getErrorMessage(r any) string {
	switch err := r.(type) {
	case error:
		return m.getErrorFromException(err).Error()

	default:
		return errutil.NewErrInternal("something went wrong, please try again later").Error()
	}
}

// getErrorFromException constructs an error message from an exception. It checks if the
// error is of type errutil.ErrInternal and if not, wraps the error into an errutil.ErrInternal.
func (m *saveExceptionHandlerImpl) getErrorFromException(err error) error {
	var errInternal *errutil.ErrInternal
	if !errors.As(err, &errInternal) {
		return errutil.NewErrInternal(err.Error())
	}

	return err
}
