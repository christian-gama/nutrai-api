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
func Error(err error) map[string]any {
	var errs *errutil.Error
	if e, ok := err.(*errutil.Error); ok {
		errs = e
	} else {
		errs = errutil.Append(errs, err)
	}

	if env.IsProduction && env.Config.Debug {
		return ErrorDebug(errs)
	}

	return map[string]any{
		"errors": errs.Errors,
	}
}

// ErrorDebug returns a JSON with the error and the stack trace. The status is always false.
func ErrorDebug(err *errutil.Error) map[string]any {
	stack := string(debug.Stack())

	return map[string]any{
		"errors": err.Errors,
		"stack":  stack,
	}
}

// Data returns a JSON with the data. The status is always true.
func Data(data any) any {
	return data
}
