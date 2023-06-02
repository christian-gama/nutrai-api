package middleware

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
	"github.com/christian-gama/nutrai-api/pkg/unit"
	"github.com/gin-gonic/gin"
)

type LimitBodySize = middleware.Middleware

func NewLimitBodySize() LimitBodySize {
	return middleware.NewMiddleware(func(ctx *gin.Context) {
		const maxBodySize = 3 * unit.Megabyte

		if ctx.Request.ContentLength > maxBodySize {
			response.BadRequest(
				ctx,
				errors.Invalid("body", "body size must be less than %d bytes", maxBodySize),
			)
			return
		}

		ctx.Next()
	})
}
