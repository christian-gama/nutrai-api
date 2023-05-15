package http

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Middlewarel is an interface that represents a middleware for gin.
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
func NewMiddleware(handler func(*gin.Context)) Middleware {
	if handler == nil {
		panic(errors.New("handler is nil"))
	}

	return &middlewareImpl{
		Handler: handler,
	}
}

// Handle implements Middleware.
func (c *middlewareImpl) Handle(ctx *gin.Context) {
	c.Handler(ctx)
}
