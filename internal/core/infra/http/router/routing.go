package router

import (
	"errors"
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/auth/api/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/pkg/slice"
	"github.com/gin-gonic/gin"
)

// Route holds the data of a route.
type Route struct {
	Middlewares []http.Middleware
	Controller  http.Controller
}

// Routing holds the data of a routing.
type Routing struct {
	Group       string
	Routes      []*Route
	Middlewares []http.Middleware
}

func (r *Routing) validate(route *Route) {
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

// Register registers the routing into the router.
func (r *Routing) Register(router *gin.RouterGroup) {
	var group *gin.RouterGroup
	if r.Group != "" {
		group = router.Group(r.Group)
	}

	for _, middleware := range r.Middlewares {
		if group != nil {
			group.Use(middleware.Handle)
		} else {
			router.Use(middleware.Handle)
		}
	}

	for _, route := range r.Routes {
		r.validate(route)

		route.Middlewares = append(route.Middlewares, route.Controller)

		handlers := slice.
			Map(route.Middlewares, func(middleware http.Middleware) gin.HandlerFunc {
				return middleware.Handle
			}).
			Build()

		path := route.Controller.Path()
		if len(route.Controller.Params()) > 0 {
			path = route.Controller.Params().ToPath(path)
		}

		// Must be unshifted because the auth middleware must be the first one.
		if !route.Controller.IsPublic() {
			handlers = slice.
				Unshift(handlers, middleware.MakeAuthHandler().Handle).
				Build()
		}

		if group != nil {
			group.Handle(route.Controller.Method().String(), path.String(), handlers...)
		} else {
			router.Handle(route.Controller.Method().String(), path.String(), handlers...)
		}
	}
}
