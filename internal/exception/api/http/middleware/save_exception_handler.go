package middleware

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/exception/app/command"
	"github.com/christian-gama/nutrai-api/pkg/errutil"
	"github.com/gin-gonic/gin"
)

// SaveExceptionHandler is the middleware to add error.
type SaveExceptionHandler = http.Middleware

// NewSaveExceptionHandler creates a new SaveException.
func NewSaveExceptionHandler(
	saveExceptionHandler command.SaveExceptionHandler,
) SaveExceptionHandler {
	if saveExceptionHandler == nil {
		panic(errors.New("command.SaveExceptionHandler cannot be nil"))
	}

	return http.NewMiddleware(
		func(ctx *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					var msg string

					if err, ok := r.(error); ok {
						var errInternal *errutil.ErrInternal
						msg = err.Error()

						if !errors.As(err, &errInternal) {
							msg = errutil.NewErrInternal(msg).Error()
						}

						saveExceptionHandler.Handle(
							context.Background(),
							&command.SaveExceptionInput{
								Message: msg,
								Stack:   string(debug.Stack()),
							},
						)
					} else {
						msg = errutil.NewErrInternal("something went wrong, please try again later").Error()

						saveExceptionHandler.Handle(
							context.Background(),
							&command.SaveExceptionInput{
								Message: msg,
								Stack:   string(debug.Stack()),
							})
					}

					ctx.AbortWithStatusJSON(
						http.StatusInternalServerError,
						http.Error(errors.New(msg)),
					)
				}
			}()

			ctx.Next()
		},
	)
}
