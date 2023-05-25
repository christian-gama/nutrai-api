package response

import (
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
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
		{ErrorAssertFunc: isErrRepository, Response: BadRequest},
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
	var errNotFound *errutil.ErrNotFound
	return errors.As(err, &errNotFound)
}

func isErrInternal(err error) bool {
	var errInternal *errutil.ErrInternal
	return errors.As(err, &errInternal)
}

func isErrInvalid(err error) bool {
	var errInvalid *errutil.ErrInvalid
	return errors.As(err, &errInvalid)
}

func isErrUnauthorized(err error) bool {
	var errUnauthorized *errutil.ErrUnauthorized
	return errors.As(err, &errUnauthorized)
}

func isErrRepository(err error) bool {
	var errRepository *errutil.ErrRepository
	return errors.As(err, &errRepository)
}
