package http

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// Error returns a JSON with the error. The status is always false.
// It will normalize the incoming error to a *errutil.Error, in case
// it's not. In a non-production environment, it will also return
// the stack trace.
func Error(err error) *ResponseBody {
	var errs *errutil.Error
	if e, ok := err.(*errutil.Error); ok {
		errs = e
	} else {
		errs = errutil.Append(errs, err)
	}

	if env.App.Env != env.Production && env.Config.Debug {
		return ErrorDebug(errs)
	}

	return &ResponseBody{
		Status: false,
		Data:   errs.Error(),
	}
}

// ErrorDebug returns a JSON with the error and the stack trace. The status is always false.
func ErrorDebug(err error) *ResponseBody {
	stack := string(debug.Stack())

	return &ResponseBody{
		Status: false,
		Data:   err.Error(),
		Stack:  stack,
	}
}

// Data returns a JSON with the data. The status is always true.
func Data(data any) *ResponseBody {
	return &ResponseBody{
		Status: true,
		Data:   data,
	}
}

// Response returns a JSON with the data based on the result of the handler.
// If the handler panics, it will return a JSON with the error according to the error type.
func Response(ctx *gin.Context, handler func()) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				var errNotFound *errutil.ErrNotFound
				if errors.As(err, &errNotFound) {
					NotFound(ctx, err)
					return
				}

				var errInternal *errutil.ErrInternal
				if errors.As(err, &errInternal) {
					InternalServerError(ctx, err)
					return
				}

				var invalidErr *errutil.ErrInvalid
				if errors.As(err, &invalidErr) {
					BadRequest(ctx, err)
					return
				}

				if e, ok := err.(*errutil.Error); ok {
					BadRequest(ctx, e)
				} else {
					InternalServerError(ctx, err)
				}
			} else {
				InternalServerError(ctx, fmt.Errorf("%v", r))
			}
		}
	}()

	handler()
}
