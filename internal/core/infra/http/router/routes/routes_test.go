package routes_test

import (
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	httpmiddleware "github.com/christian-gama/nutrai-api/internal/core/infra/http/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	routesMiddleware "github.com/christian-gama/nutrai-api/internal/core/infra/http/router/middleware"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router/routes"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/gin-gonic/gin"
)

type RoutesSuite struct {
	suite.Suite
}

func TestRoutesSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(RoutesSuite))
}

func (s *RoutesSuite) TestApi() {
	routesMiddleware.SetAuthMiddleware(MakeAuthMiddleware())

	setupRouter := func() {
		gin.SetMode(gin.TestMode)
		router.Router = gin.New()
	}

	s.Run("should register a route within the api group", func() {
		setupRouter()

		routes.Api().SetController(&Controller{
			method:   http.MethodGet,
			path:     "/",
			handler:  func(c *gin.Context) string { return "ok" },
			isPublic: true,
		})

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusOK, response.Code)
	})

	s.Run("should register a route within the api group with a middleware", func() {
		setupRouter()

		routes.Api().
			SetMiddleware(MakeMiddleware("middleware", "ok")).
			SetController(&Controller{
				method:   http.MethodGet,
				path:     "/",
				handler:  func(c *gin.Context) string { return c.GetString("middleware") },
				isPublic: true,
			})

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("ok", response.Body.String())
	})

	s.Run("should not use a middleware that was registered after the controller", func() {
		setupRouter()

		routes.Api().
			SetMiddleware(MakeMiddleware("middleware", "first_middleware")).
			SetController(&Controller{
				method:   http.MethodGet,
				path:     "/",
				isPublic: true,
				handler:  func(c *gin.Context) string { return c.GetString("middleware") },
			}).
			SetMiddleware(MakeMiddleware("middleware", "second_middleware"))

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("first_middleware", response.Body.String())
	})

	s.Run("should register local middlewares", func() {
		setupRouter()

		routes.Api().
			SetController(&Controller{
				method:   http.MethodGet,
				isPublic: true,
				path:     "/",
				handler:  func(c *gin.Context) string { return c.GetString("middleware") },
			},
				MakeMiddleware("middleware", "local_middleware"),
			)

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("local_middleware", response.Body.String())
	})

	s.Run("should register local middlewares in the correct order", func() {
		setupRouter()

		routes.Api().
			SetController(&Controller{
				method:   http.MethodGet,
				isPublic: true,
				path:     "/",
				handler:  func(c *gin.Context) string { return c.GetString("middleware") },
			},
				MakeMiddleware("middleware", "first_local_middleware"),
				MakeMiddleware("middleware", "second_local_middleware"),
			)

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("second_local_middleware", response.Body.String())
	})

	s.Run("should register with a new group", func() {
		setupRouter()

		routes.Api("new").
			SetController(&Controller{
				method:   http.MethodGet,
				isPublic: true,
				path:     "/",
				handler:  func(c *gin.Context) string { return "ok" },
			})

		response := Request(http.MethodGet, "/api/new/")

		s.Equal(http.StatusOK, response.Code)
	})

	s.Run("should register with a new group with a middleware", func() {
		setupRouter()

		routes.Api("new").
			SetMiddleware(MakeMiddleware("middleware", "ok")).
			SetController(&Controller{
				method:   http.MethodGet,
				path:     "/",
				isPublic: true,
				handler:  func(c *gin.Context) string { return c.GetString("middleware") },
			})

		response := Request(http.MethodGet, "/api/new/")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("ok", response.Body.String())
	})

	s.Run("should register a route with params", func() {
		setupRouter()

		routes.Api().
			SetController(&Controller{
				method:   http.MethodGet,
				path:     "",
				params:   controller.AddParams("id"),
				isPublic: true,
				handler:  func(c *gin.Context) string { return c.Param("id") },
			})

		response := Request(http.MethodGet, "/api/1")

		s.Equal(http.StatusOK, response.Code)
		s.Equal("1", response.Body.String())
	})

	s.Run("should register a private route", func() {
		setupRouter()

		routes.Api().
			SetController(&Controller{
				method:  http.MethodGet,
				path:    "/",
				handler: func(c *gin.Context) string { return "ok" },
			})

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusUnauthorized, response.Code)
	})

	s.Run("should call auth middleware before any other middleware", func() {
		setupRouter()
		called := false

		routes.Api().
			SetController(&Controller{
				method:  http.MethodGet,
				path:    "/",
				handler: func(c *gin.Context) string { return "ok" },
			},
				httpmiddleware.NewMiddleware(func(ctx *gin.Context) {
					called = true
				}),
			)

		response := Request(http.MethodGet, "/api/")

		s.Equal(http.StatusUnauthorized, response.Code)
		s.False(
			called,
			"should not call the middleware, because the chain should be broken by the auth middleware",
		)
	})
}

func Request(
	method http.Method,
	path string,
) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()

	router.Router.ServeHTTP(response, httptest.NewRequest(string(method), path, nil))

	return response
}

type Controller struct {
	status    int
	method    http.Method
	path      controller.Path
	params    controller.Params
	isPublic  bool
	rateLimit int
	handler   func(*gin.Context) string
}

func (h *Controller) Handle(c *gin.Context) {
	c.String(h.status, h.handler(c))
}

func (h *Controller) Method() http.Method {
	return h.method
}

func (h *Controller) Path() controller.Path {
	return h.path
}

func (h *Controller) IsPublic() bool {
	return h.isPublic
}

func (h *Controller) Params() controller.Params {
	return h.params
}

func (h *Controller) RPM() int {
	return h.rateLimit
}

type Middleware struct {
	Value string
	Key   string
}

func (m *Middleware) Handle(c *gin.Context) {
	c.Set(m.Key, m.Value)
	c.Next()
}

func MakeMiddleware(key, value string) *Middleware {
	return &Middleware{
		Value: value,
		Key:   key,
	}
}

type AuthMiddleware struct{}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func MakeAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
