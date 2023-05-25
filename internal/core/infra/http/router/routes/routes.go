package routes

import (
	authMiddleware "github.com/christian-gama/nutrai-api/internal/auth/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/gin-gonic/gin"
)

func SetGlobalMiddleware(middleware ...middleware.Middleware) {
	router.Router.Use(func(c *gin.Context) {
		for _, m := range middleware {
			m.Handle(c)
		}
	})
}

type routes struct {
	group *gin.RouterGroup
}

// Api initializes a new instance of routes with a given group.
func Api(group ...string) *routes {
	return &routes{
		group: router.Router.Group("api").Group(slice.FirstElementOrDefault(group)),
	}
}

// SetMiddleware sets the middleware to the router group. This middlewares runs before any
// controller and auth middleware.
func (r *routes) SetMiddleware(middleware middleware.Middleware) *routes {
	r.group.Use(middleware.Handle)
	return r
}

// SetController adds a controller with optional middleware to the router group.
func (r *routes) SetController(
	controller controller.Controller,
	middleware ...middleware.Middleware,
) *routes {
	handlers := r.initializeHandlers(controller, middleware...)
	path := r.buildPath(controller)
	r.registerHandlersToGroup(controller, path, handlers)
	return r
}

// initializeHandlers is a helper function that initializes handlers with controller and middleware.
func (r *routes) initializeHandlers(
	controller controller.Controller,
	middleware ...middleware.Middleware,
) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, len(middleware)+1)
	for i, middleware := range middleware {
		handlers[i] = middleware.Handle
	}
	handlers[len(handlers)-1] = controller.Handle

	handlers = r.addAuthIfNeeded(controller, handlers)

	return handlers
}

// buildPath is a helper function that builds the path for the controller.
func (r *routes) buildPath(controller controller.Controller) string {
	path := controller.Path()
	if len(controller.Params()) > 0 {
		path = controller.Params().ToPath(path)
	}

	return path.String()
}

// addAuthIfNeeded is a helper function that adds an auth middleware if the controller is not
// public.
func (r *routes) addAuthIfNeeded(
	controller controller.Controller,
	handlers []gin.HandlerFunc,
) []gin.HandlerFunc {
	if !controller.IsPublic() {
		handlers = slice.Unshift(handlers, authMiddleware.MakeAuth().Handle).Build()
	}

	return handlers
}

// registerHandlersToGroup is a helper function that adds handlers to the router group.
func (r *routes) registerHandlersToGroup(
	controller controller.Controller,
	path string,
	handlers []gin.HandlerFunc,
) {
	r.group.Handle(controller.Method().String(), path, handlers...)
}
