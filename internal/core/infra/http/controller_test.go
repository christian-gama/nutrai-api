package http_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http"
	"github.com/christian-gama/nutrai-api/testutils/gintest"
	"github.com/christian-gama/nutrai-api/testutils/suite"
	"github.com/gin-gonic/gin"
)

type ControllerSuite struct {
	suite.Suite
}

func TestControllerSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ControllerSuite))
}

func (s *ControllerSuite) TestNewController() {
	makeSut := func() (
		func(handler http.Handler[any], opts http.ControllerOptions) http.Controller,
		func(*gin.Context, *any),
		http.ControllerOptions,
	) {
		handler := func(*gin.Context, *any) {}
		opts := http.ControllerOptions{
			Method:   http.MethodGet,
			Path:     http.JoinPath(""),
			IsPublic: true,
			Params:   http.AddParams("id"),
		}
		sut := http.NewController[any]

		return sut, handler, opts
	}

	s.Run("do not panic when all args are valid", func() {
		sut, handler, opts := makeSut()

		s.NotPanics(func() { sut(handler, opts) })
	})

	s.Run("panic when handler is nil", func() {
		sut, _, opts := makeSut()

		s.Panics(func() { sut(nil, opts) })
	})

	s.Run("panic when method is empty", func() {
		sut, handler, opts := makeSut()

		opts.Method = ""

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when path is empty", func() {
		sut, handler, opts := makeSut()

		opts.Path = ""

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when params are invalid", func() {
		sut, handler, opts := makeSut()

		opts.Params = http.Params{"invalid param"}

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when params are not unique", func() {
		sut, handler, opts := makeSut()

		opts.Params = http.AddParams("id").Add("id")

		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when path is invalid", func() {
		sut, handler, opts := makeSut()

		opts.Path = "invalid path"
		s.Panics(func() { sut(handler, opts) })

		opts.Path = "/invalid/:path"
		s.Panics(func() { sut(handler, opts) })
	})

	s.Run("panic when method is invalid", func() {
		sut, handler, opts := makeSut()

		opts.Method = "invalid method"

		s.Panics(func() { sut(handler, opts) })
	})
}

func (s *ControllerSuite) TestController() {
	sut := http.NewController(
		func(*gin.Context, *any) {},

		http.ControllerOptions{
			Method:   http.MethodPut,
			Path:     http.JoinPath(""),
			IsPublic: true,
			Params:   http.AddParams("id"),
		},
	)

	s.Run("controller.Method returns the correct method", func() {
		sut := sut.Method

		method := sut()

		s.EqualValues(http.MethodPut, method)
	})

	s.Run("controller.Path returns the correct path", func() {
		sut := sut.Path

		path := sut()

		s.EqualValues("/", path)
	})

	s.Run("controller.IsPublic returns the correct value", func() {
		sut := sut.IsPublic

		isPublic := sut()

		s.True(isPublic)
	})

	s.Run("controller.Params returns the correct params", func() {
		sut := sut.Params

		params := sut()

		s.Equal([]string{"id"}, params.Slice())
	})
}

type Payload struct {
	Name string `json:"name" validate:"max=10"`
	ID   int    `uri:"id"    validate:"gte=1"`
	Age  int    `form:"age"  validate:"lte=100"`
}

func (s *ControllerSuite) TestHandle() {
	makeSut := func() (http.Controller, *Payload) {
		payload := &Payload{
			Name: "John Doe",
			ID:   1,
			Age:  20,
		}

		handler := func(c *gin.Context, p *Payload) {
			s.Equal(payload.Name, p.Name)
			s.Equal(payload.ID, p.ID)
			s.Equal(payload.Age, p.Age)
		}

		controller := http.NewController(
			handler,
			http.ControllerOptions{
				Method:   http.MethodPut,
				Path:     http.JoinPath(""),
				IsPublic: true,
				Params:   http.AddParams("id"),
			},
		)

		return controller, payload
	}

	s.Run("returns 200 when all params are valid", func() {
		sut, payload := makeSut()

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": payload.Name},
			Params:  []string{fmt.Sprintf("%d", payload.ID)},
			Queries: fmt.Sprintf("age=%d", payload.Age),
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})

	s.Run("returns 400 when params are invalid", func() {
		sut, payload := makeSut()
		payload.Name = "this is an invalid name"

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": payload.Name},
			Params:  []string{fmt.Sprintf("%d", payload.ID)},
			Queries: fmt.Sprintf("age=%d", payload.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 400 when params are not valid", func() {
		sut, payload := makeSut()
		payload.ID = -1

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": payload.Name},
			Params:  []string{fmt.Sprintf("%d", payload.ID)},
			Queries: fmt.Sprintf("age=%d", payload.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 400 when query is not valid", func() {
		sut, payload := makeSut()
		payload.Age = 101

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:    map[string]any{"name": payload.Name},
			Params:  []string{fmt.Sprintf("%d", payload.ID)},
			Queries: fmt.Sprintf("age=%d", payload.Age),
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("returns 200 even when query is not provided", func() {
		sut, payload := makeSut()
		payload.Age = 0

		ctx := gintest.MustRequest(sut, gintest.Option{
			Data:   map[string]any{"name": payload.Name},
			Params: []string{fmt.Sprintf("%d", payload.ID)},
		})

		s.Equal(http.StatusOK, ctx.Writer.Status())
	})
}
