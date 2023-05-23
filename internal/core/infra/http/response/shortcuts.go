package response

import (
	"runtime/debug"

	"github.com/christian-gama/nutrai-api/config/env"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
)

// Error returns a JSON with the error. The status is always false.
// It will normalize the incoming error to a *errutil.Error, in case
// it's not. In a non-production environment, it will also return
// the stack trace.
func Error(err error) *Body {
	var errs *errutil.Error
	if e, ok := err.(*errutil.Error); ok {
		errs = e
	} else {
		errs = errutil.Append(errs, err)
	}

	if env.App.Env != env.Production && env.Config.Debug {
		return ErrorDebug(errs)
	}

	return &Body{
		Status: false,
		Data:   errs.Error(),
	}
}

// ErrorDebug returns a JSON with the error and the stack trace. The status is always false.
func ErrorDebug(err error) *Body {
	stack := string(debug.Stack())

	return &Body{
		Status: false,
		Data:   err.Error(),
		Stack:  stack,
	}
}

// Data returns a JSON with the data. The status is always true.
func Data(data any) *Body {
	return &Body{
		Status: true,
		Data:   data,
	}
}
