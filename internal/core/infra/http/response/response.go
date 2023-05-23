package response

import (
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

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

				var errInvalid *errutil.ErrInvalid
				if errors.As(err, &errInvalid) {
					BadRequest(ctx, err)
					return
				}

				var errUnauthorized *errutil.ErrUnauthorized
				if errors.As(err, &errUnauthorized) {
					Unauthorized(ctx, err)
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
