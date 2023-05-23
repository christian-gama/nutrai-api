package middleware

import (
	"errors"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/gin-gonic/gin"
)

type Handler func(*gin.Context)

// Middleware is an interface that represents a middleware for gin.
// It extracts the body, query and params from the request and passes it to the handler.
type Middleware interface {
	// Handle is the function that will be called by the router.
	Handle(ctx *gin.Context)
}

// middlewareImpl is the implementation of the Middleware interface.
type middlewareImpl struct {
	Handler func(*gin.Context)
}

// NewMiddleware creates a new middleware.
func NewMiddleware(handler Handler) Middleware {
	if handler == nil {
		panic(errors.New("handler is nil"))
	}

	return &middlewareImpl{
		Handler: handler,
	}
}

// Handle implements Middleware.
func (c *middlewareImpl) Handle(ctx *gin.Context) {
	response.Response(ctx, func() { c.Handler(ctx) })
}
