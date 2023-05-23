package router

import (
	"errors"
	"fmt"

	authMiddleware "github.com/christian-gama/nutrai-api/internal/auth/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/gin-gonic/gin"
)

// Route represents a single HTTP route in the application. Each route consists of one or more
// Middlewares that process the request, and a Controller that generates the response.
type Route struct {
	Middlewares []middleware.Middleware
	Controller  controller.Controller
}

// validate ensures that a Route's Controller is correctly configured. If not,
// it panics with a relevant error message. It checks if Controller is not nil,
// and whether its Method and Path are non-empty.
func (route *Route) validate() {
	if route.Controller == nil {
		panic(errors.New("controller is nil"))
	}
	if route.Controller.Method() == "" {
		panic(fmt.Errorf("method is empty for controller %v", route.Controller))
	}
	if route.Controller.Path() == "" {
		panic(fmt.Errorf("path is empty for controller %v", route.Controller))
	}
}

// appendControllerToMiddlewares adds the Route's Controller to it's list of Middlewares.
func (route *Route) appendControllerToMiddlewares() {
	route.Middlewares = append(route.Middlewares, route.Controller)
}

// generateHandlersFromMiddlewares creates a slice of gin.HandlerFunc from the Route's Middlewares.
func (route *Route) generateHandlersFromMiddlewares() []gin.HandlerFunc {
	return slice.Map(route.Middlewares, func(middleware middleware.Middleware) gin.HandlerFunc {
		return middleware.Handle
	}).Build()
}

// generatePathFromController generates the Route's path string from its Controller.
// If the Controller has Params, it generates a path string that includes those Params.
func (route *Route) generatePathFromController() string {
	path := route.Controller.Path()
	if len(route.Controller.Params()) > 0 {
		path = route.Controller.Params().ToPath(path)
	}
	return path.String()
}

// addAuthHandlerToHandlers adds an Auth middleware to the start of the handlers slice,
// if the Route's Controller is not marked as Public. This ensures that authentication
// is performed before any other middleware or controller action for non-public routes.
func (route *Route) addAuthHandlerToHandlers(handlers []gin.HandlerFunc) []gin.HandlerFunc {
	return slice.Unshift(handlers, authMiddleware.MakeAuth().Handle).Build()
}

// Routing encapsulates the details of a group of related HTTP routes. It includes a Group string
// that may be used as a URL prefix for all Routes, and a slice of Middlewares that are applied to
// all Routes.
type Routing struct {
	Group       string
	Routes      []*Route
	Middlewares []middleware.Middleware
}

// Register takes a gin.RouterGroup and registers all of Routing's Routes on it.
// It first applies all of Routing's Middlewares, then registers each Route one by one.
// If Routing.Group is non-empty, it creates a new RouterGroup with the group as the URL prefix.
func (r *Routing) Register(router *gin.RouterGroup) {
	group := r.defineRouterGroup(router)

	r.addGroupMiddlewares(router, group)
	r.registerRoutes(router, group)
}

// defineRouterGroup creates a new gin.RouterGroup if Routing.Group is non-empty, otherwise it
// returns nil.
func (r *Routing) defineRouterGroup(router *gin.RouterGroup) *gin.RouterGroup {
	if r.Group != "" {
		return router.Group(r.Group)
	}
	return nil
}

// addGroupMiddlewares applies all of Routing's Middlewares to the provided gin.RouterGroup. If a
// group is provided, it applies the middlewares to the group, otherwise it applies them to the
// router.
func (r *Routing) addGroupMiddlewares(router *gin.RouterGroup, group *gin.RouterGroup) {
	for _, middleware := range r.Middlewares {
		if group != nil {
			group.Use(middleware.Handle)
		} else {
			router.Use(middleware.Handle)
		}
	}
}

// registerRoutes registers all of Routing's Routes onto the provided gin.RouterGroup. If a group is
// provided, it registers the routes on the group, otherwise it registers them on the
// router.
func (r *Routing) registerRoutes(router *gin.RouterGroup, group *gin.RouterGroup) {
	for _, route := range r.Routes {
		route.validate()

		route.appendControllerToMiddlewares()

		handlers := route.generateHandlersFromMiddlewares()

		path := route.generatePathFromController()

		if !route.Controller.IsPublic() {
			handlers = route.addAuthHandlerToHandlers(handlers)
		}

		r.addHandlerToGroupOrRouter(route, router, group, handlers, path)
	}
}

// addHandlerToGroupOrRouter adds the handlers of a Route to the provided gin.RouterGroup. If a
// group is provided, it adds the handlers to the group, otherwise it adds them to the router.
func (r *Routing) addHandlerToGroupOrRouter(
	route *Route,
	router *gin.RouterGroup,
	group *gin.RouterGroup,
	handlers []gin.HandlerFunc,
	path string,
) {
	if group != nil {
		group.Handle(route.Controller.Method().String(), path, handlers...)
	} else {
		router.Handle(route.Controller.Method().String(), path, handlers...)
	}
}
