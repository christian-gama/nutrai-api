package response

import (
	_errors "errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	ErrorAssertFunc func(error) bool
	Response        func(*gin.Context, error)
}

// Response returns a JSON with the data based on the result of the handler.
// If the handler panics, it will return a JSON with the error according to the error type.
func Response(ctx *gin.Context, handler func()) {
	defer handleRecovery(ctx)
	handler()
}

func handleRecovery(ctx *gin.Context) {
	if r := recover(); r != nil {
		handleError(ctx, r)
	}
}

func handleError(ctx *gin.Context, r any) {
	if err, ok := r.(error); ok {
		if !handleSpecificErrors(ctx, err) {
			handleGenericErrors(ctx, err)
		}
		return
	}

	InternalServerError(ctx, fmt.Errorf("%v", r))
}

func handleSpecificErrors(ctx *gin.Context, err error) bool {
	errorResponses := []errorResponse{
		{ErrorAssertFunc: isErrAlreadyExists, Response: Conflict},
		{ErrorAssertFunc: isErrInternalServerError, Response: InternalServerError},
		{ErrorAssertFunc: isErrInvalid, Response: BadRequest},
		{ErrorAssertFunc: isErrNoChanges, Response: BadRequest},
		{ErrorAssertFunc: isErrNotFound, Response: NotFound},
		{ErrorAssertFunc: isErrRequired, Response: BadRequest},
		{ErrorAssertFunc: isErrTimeout, Response: GatewayTimeout},
		{ErrorAssertFunc: isErrTooManyRequests, Response: TooManyRequests},
		{ErrorAssertFunc: isErrUnauthorized, Response: Unauthorized},
		{ErrorAssertFunc: isErrUnavailable, Response: ServiceUnavailable},
	}

	for _, errorHandler := range errorResponses {
		if errorHandler.ErrorAssertFunc(err) {
			errorHandler.Response(ctx, err)
			return true
		}
	}
	return false
}

func handleGenericErrors(ctx *gin.Context, err error) {
	if e, ok := err.(*errutil.Error); ok {
		BadRequest(ctx, e)
		return
	}

	InternalServerError(ctx, fmt.Errorf("%v", err))
}

func isErrAlreadyExists(err error) bool {
	var errAlreadyExists *errors.ErrAlreadyExists
	return _errors.As(err, &errAlreadyExists)
}

func isErrInternalServerError(err error) bool {
	var errInternalServerError *errors.ErrInternalServerError
	return _errors.As(err, &errInternalServerError)
}

func isErrInvalid(err error) bool {
	var errInvalid *errors.ErrInvalid
	return _errors.As(err, &errInvalid)
}

func isErrNoChanges(err error) bool {
	var errNoChanges *errors.ErrNoChanges
	return _errors.As(err, &errNoChanges)
}

func isErrNotFound(err error) bool {
	var errNotFound *errors.ErrNotFound
	return _errors.As(err, &errNotFound)
}

func isErrRequired(err error) bool {
	var errRequired *errors.ErrRequired
	return _errors.As(err, &errRequired)
}

func isErrTimeout(err error) bool {
	var errTimeout *errors.ErrTimeout
	return _errors.As(err, &errTimeout)
}

func isErrTooManyRequests(err error) bool {
	var errTooManyRequests *errors.ErrTooManyRequests
	return _errors.As(err, &errTooManyRequests)
}

func isErrUnauthorized(err error) bool {
	var errUnauthorized *errors.ErrUnauthorized
	return _errors.As(err, &errUnauthorized)
}

func isErrUnavailable(err error) bool {
	var errUnavailable *errors.ErrUnavailable
	return _errors.As(err, &errUnavailable)
}
