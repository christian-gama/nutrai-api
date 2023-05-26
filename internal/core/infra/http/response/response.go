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
		{ErrorAssertFunc: isErrNotFound, Response: NotFound},
		{ErrorAssertFunc: isErrInternal, Response: InternalServerError},
		{ErrorAssertFunc: isErrInvalid, Response: BadRequest},
		{ErrorAssertFunc: isErrUnauthorized, Response: Unauthorized},
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

func isErrNotFound(err error) bool {
	var errNotFound *errors.ErrNotFound
	return _errors.As(err, &errNotFound)
}

func isErrInternal(err error) bool {
	var errInternal *errors.ErrInternalServerError
	return _errors.As(err, &errInternal)
}

func isErrInvalid(err error) bool {
	var errInvalid *errors.ErrInvalid
	return _errors.As(err, &errInvalid)
}

func isErrUnauthorized(err error) bool {
	var errUnauthorized *errors.ErrUnauthorized
	return _errors.As(err, &errUnauthorized)
}
