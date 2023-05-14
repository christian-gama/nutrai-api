package routing_test

import (
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/http"
	"github.com/christian-gama/nutrai-api/internal/shared/infra/router/routing"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/gin-gonic/gin"
)

type RoutingSuite struct {
	suite.Suite
}

func TestRoutingSuite(t *testing.T) {
	suite.RunIntegrationTest(t, new(RoutingSuite))
}

type Controller struct {
	status  int
	method  http.Method
	path    http.Path
	params  http.Params
	handler func(*gin.Context) string
}

func (h *Controller) Handle(c *gin.Context) {
	c.String(h.status, h.handler(c))
}

func (h *Controller) Method() http.Method {
	return h.method
}

func (h *Controller) Path() http.Path {
	return h.path
}

func (h *Controller) IsPublic() bool {
	return true
}

func (h *Controller) Params() http.Params {
	return h.params
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

func (s *RoutingSuite) TestRegister() {
	s.Run("create routes successfully", func() {
		routes := &routing.Routing{
			Group: "",
			Routes: []*routing.Route{
				{
					Controller: &Controller{
						status:  http.StatusOK,
						handler: func(ctx *gin.Context) string { return "Hello, World!" },
						method:  http.MethodGet,
						path:    http.JoinPath(""),
					},
				},
			},
		}

		w := Register(routes, http.MethodGet, "/test/")

		s.Equal(http.StatusOK, w.Code)
	})

	s.Run("create routes with group successfully", func() {
		routes := &routing.Routing{
			Group: "/api",
			Routes: []*routing.Route{
				{
					Controller: &Controller{
						status:  http.StatusOK,
						handler: func(ctx *gin.Context) string { return "Hello, World!" },
						method:  http.MethodGet,
						path:    "/",
						params:  http.AddParams("id"),
					},
				},
			},
		}

		w := Register(routes, http.MethodGet, "/test/api/1")

		s.Equal(http.StatusOK, w.Code)
	})

	s.Run("panics when controller has missing attributes", func() {
		routes := &routing.Routing{
			Group: "",
			Routes: []*routing.Route{
				{
					Controller: &Controller{
						status: http.StatusOK,
						handler: func(ctx *gin.Context) string {
							return "Hello, World!"
						},
					},
				},
			},
		}

		s.Panics(func() {
			Register(routes, http.MethodGet, "/test/")
		})
	})

	s.Run("works with middleware correctly", func() {
		routes := &routing.Routing{
			Routes: []*routing.Route{
				{
					Middlewares: []http.Middleware{
						MakeMiddleware("test", "success"),
					},
					Controller: &Controller{
						status:  http.StatusOK,
						handler: func(ctx *gin.Context) string { return ctx.GetString("test") },
						method:  http.MethodGet,
						path:    "/",
					},
				},
			},
		}

		w := Register(routes, http.MethodGet, "/test/")

		s.Equal(http.StatusOK, w.Code)
		s.Equal("success", w.Body.String())
	})

	s.Run("returns 404 when route is not found", func() {
		routes := &routing.Routing{
			Group: "",
			Routes: []*routing.Route{
				{
					Controller: &Controller{
						status:  http.StatusOK,
						handler: func(ctx *gin.Context) string { return "Hello, World!" },
						method:  http.MethodGet,
						path:    "/",
					},
				},
			},
		}

		w := Register(routes, http.MethodGet, "/not-found")

		s.Equal(http.StatusNotFound, w.Code)
	})

	s.Run("works with global middleware correctly", func() {
		routes := &routing.Routing{
			Middlewares: []http.Middleware{
				MakeMiddleware("test", "success"),
			},
			Routes: []*routing.Route{
				{
					Controller: &Controller{
						status:  http.StatusOK,
						handler: func(ctx *gin.Context) string { return ctx.GetString("test") },
						method:  http.MethodGet,
						path:    "/",
					},
				},
			},
		}

		w := Register(routes, http.MethodGet, "/test/")

		s.Equal(http.StatusOK, w.Code)
		s.Equal("success", w.Body.String())
	})
}

func Register(routing *routing.Routing, method http.Method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	router := gin.New()
	v1 := router.Group("/test")
	routing.Register(v1)

	router.ServeHTTP(w, httptest.NewRequest(string(method), path, nil))

	return w
}
